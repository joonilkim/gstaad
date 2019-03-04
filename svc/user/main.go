package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	_log "log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	mw "gstaad/pkg/middleware"
	pb "gstaad/pkg/proto/user"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	Version = "0.0.0+1"
)

var (
	port     = ":" + getenv("PORT", "9000")
	grpcPort = ":" + getenv("GRPC_PORT", "9001")

	log  *logrus.Logger
	logw io.WriteCloser
)

func init() {
	if os.Getenv("APP_ENV") == "production" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	log = logrus.New()
	logw = log.Writer()
}

func restServer(c context.Context, upstream string) http.Handler {
	op := []grpc.DialOption{grpc.WithInsecure()}
	gw := runtime.NewServeMux()
	must(pb.RegisterUserServiceHandlerFromEndpoint(c, gw, upstream, op))

	sv := http.NewServeMux()
	sv.HandleFunc("/ping", servePing)
	sv.Handle("/", gw)
	return sv
}

func grpcServer(cc *connectors) (sv *grpc.Server) {
	sv = grpc.NewServer()

	s := &server{cc}
	pb.RegisterUserServiceServer(sv, s)
	reflection.Register(sv)
	return
}

func servePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func startGrpc(addr string) *grpc.Server {
	cc, er := NewConnectors()
	must(er)

	gconn, er := net.Listen("tcp", addr)
	must(er)
	gs := grpcServer(cc)

	go func() {
		log.Infof("listening grpc %s", addr)
		er := gs.Serve(gconn)
		if er != nil && er != http.ErrServerClosed {
			log.Fatalf("Failed: %s\n", er)
		}
	}()
	return gs
}

func startRest(c context.Context, addr string, upstream string) *http.Server {
	l := _log.New(logw, "", 0)

	nextID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	h := restServer(c, upstream)
	h = mw.Tracing(nextID)(mw.Logging(l)(h))
	h = h2c.NewHandler(h, &http2.Server{})
	sv := &http.Server{
		Addr:           addr,
		Handler:        h,
		ErrorLog:       l,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Infof("listening rest %s", addr)
		er := sv.ListenAndServe()
		if er != nil && er != http.ErrServerClosed {
			log.Fatalf("Failed: %s\n", er)
		}
	}()
	return sv
}

func main() {
	flag.Parse()

	defer logw.Close()

	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	gs := startGrpc(grpcPort)
	sv := startRest(c, port, grpcPort)

	stopGraceful(c, gs, sv)
}

func stopGraceful(c context.Context, gs *grpc.Server, sv *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	log.Info("Stopping operation...")

	sv.Shutdown(c)
	gs.GracefulStop()
}

func must(er error) {
	if er != nil {
		panic(er)
	}
}

func getenv(key, defaultval string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultval
}

package main

import (
	"context"
	"flag"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	pb "gstaad/pkg/proto/post"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	Version = "0.0.0+1"
)

var (
	port     = ":" + getenv("PORT", "9000")
	grpcPort = ":" + getenv("GRPC_PORT", "9001")
)

func init() {
	if os.Getenv("APP_ENV") == "production" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}
}

func restServer(c context.Context, upstream string) *http.ServeMux {
	op := []grpc.DialOption{grpc.WithInsecure()}
	gw := runtime.NewServeMux()
	must(pb.RegisterPostServiceHandlerFromEndpoint(c, gw, upstream, op))

	sv := http.NewServeMux()
	sv.HandleFunc("/ping", servePing)
	sv.Handle("/", gw)
	return sv
}

func grpcServer() *grpc.Server {
	sv := grpc.NewServer()
	pb.RegisterPostServiceServer(sv, &server{})
	reflection.Register(sv)
	return sv
}

func servePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func startGrpc(addr string) *grpc.Server {
	transport := "tcp"
	if strings.HasPrefix(addr, "unix://") {
		transport = "unix"
		addr = addr[7:]
	}

	gconn, er := net.Listen(transport, addr)
	must(er)
	gsv := grpcServer()

	go func() {
		log.Infof("listening grpc %s", addr)
		er := gsv.Serve(gconn)
		if er != nil && er != http.ErrServerClosed {
			log.Fatalf("Failed: %s\n", er)
		}
	}()
	return gsv
}

func startRest(c context.Context, addr string, upstream string) *http.Server {
	h := restServer(c, upstream)
	sv := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(h, &http2.Server{}),
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

	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	gsv := startGrpc(grpcPort)
	sv := startRest(c, port, grpcPort)

	stopGraceful(c, gsv, sv)
}

func stopGraceful(c context.Context, gsv *grpc.Server, sv *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	log.Info("Stopping operation...")

	sv.Shutdown(c)
	gsv.GracefulStop()
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

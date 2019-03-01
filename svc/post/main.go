package main

import (
	"context"
	"flag"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"

	pb "gstaad/pkg/proto/post"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port     = ":9000"
	grpcPort = ":9001"
)

var (
	swaggerDir = flag.String("swagger_dir", "template", "a path to dir contains swaggers")
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
	must(pb.RegisterPostHandlerFromEndpoint(c, gw, upstream, op))

	sv := http.NewServeMux()
	sv.HandleFunc("/ping", servePing)
	sv.HandleFunc("/swagger", serveSwagger)
	sv.Handle("/", gw)
	return sv
}

func grpcServer() *grpc.Server {
	sv := grpc.NewServer()
	pb.RegisterPostServer(sv, &server{})
	reflection.Register(sv)
	return sv
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(*swaggerDir, p)
	http.ServeFile(w, r, p)
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
	conn, er := net.Listen("tcp", addr)
	must(er)
	mux := restServer(c, upstream)
	sv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		log.Infof("listening rest %s", addr)
		er := sv.Serve(conn)
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

	log.Info("Shutting down...")

	c, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()
	sv.Shutdown(c)
	gsv.GracefulStop()

	select {
	case <-c.Done():
	}
}

func must(er error) {
	if er != nil {
		panic(er)
	}
}

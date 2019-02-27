package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"

	pb "gstaad/pkg/proto/post"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":9000"
)

type server struct{}

func (s *server) Create(ctx context.Context, in *pb.PostRequest) (*pb.PostReply, error) {
	return &pb.PostReply{Message: "Hello " + in.Name}, nil
}

func restServer(ctx context.Context) (s *http.ServeMux, er error) {
	op := []grpc.DialOption{grpc.WithInsecure()}
	rest := runtime.NewServeMux()
	er = pb.RegisterPostHandlerFromEndpoint(ctx, rest, port, op)
	if er != nil {
		return
	}

	s = http.NewServeMux()
	s.Handle("/", rest)
	s.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	return
}

func grpcServer(ctx context.Context) (s *grpc.Server, er error) {
	s = grpc.NewServer()
	pb.RegisterPostServer(s, &server{})

	reflection.Register(s)
	return
}

func startServer(rest *http.ServeMux, gsrv *grpc.Server) (s *http.Server, er error) {
	conn, er := net.Listen("tcp", port)
	if er != nil {
		return
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			gsrv.ServeHTTP(w, r)
		} else {
			rest.ServeHTTP(w, r)
		}
	})

	s = &http.Server{
		Addr:    fmt.Sprintf("localhost%s", port),
		Handler: handler,
	}

	fmt.Printf("starting GRPC and REST on: %v\n", port)
	er = s.Serve(conn)
	return
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rest, er := restServer(ctx)
	must(er)
	gsrv, er := grpcServer(ctx)
	must(er)

	_, er = startServer(rest, gsrv)
	must(er)
}

func must(er error) {
	if er != nil {
		panic(er)
	}
}

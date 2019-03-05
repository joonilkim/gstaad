package main

import (
	stub "gstaad/pkg/proto/post"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	health "google.golang.org/grpc/health"
	healthstub "google.golang.org/grpc/health/grpc_health_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	*health.Server
	addr string
	grpc *grpc.Server
}

func newServer(addr string) (s *server) {
	gs := grpc.NewServer()
	s = &server{health.NewServer(), addr, gs}
	stub.RegisterPostServiceServer(gs, s)
	healthstub.RegisterHealthServer(gs, s)
	reflection.Register(gs)
	return
}

func (s *server) run() {
	cc, er := net.Listen("tcp", s.addr)
	if er != nil {
		panic(er)
	}

	go func() {
		logger.Infof("listening grpc %s", s.addr)
		er := s.grpc.Serve(cc)
		if er != nil && er != http.ErrServerClosed {
			logger.Fatalf("Failed: %s\n", er)
		}
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	logger.Info("Stopping operation...")
	s.grpc.GracefulStop()
}

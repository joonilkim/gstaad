package main

import (
	stub "gstaad/pkg/proto/user"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	Addr       string
	Server     *grpc.Server
	connectors *connectors
}

func newServer(addr string, cc *connectors) (s *server) {
	gs := grpc.NewServer()
	s = &server{addr, gs, cc}
	stub.RegisterUserServiceServer(gs, s)
	reflection.Register(gs)
	return
}

func (s *server) run() {
	cc, er := net.Listen("tcp", s.Addr)
	if er != nil {
		panic(er)
	}

	go func() {
		logger.Infof("listening grpc %s", s.Addr)
		er := s.Server.Serve(cc)
		if er != nil && er != http.ErrServerClosed {
			logger.Fatalf("Failed: %s\n", er)
		}
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	logger.Info("Stopping operation...")
	s.Server.GracefulStop()
}

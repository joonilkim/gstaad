package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	userpb "gstaad/src/userservice/pb"

	_ "google.golang.org/grpc/encoding/gzip" // Install the gzip compressor
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type server struct {
	addr string
	grpc *grpc.Server
	cc   *connectors
}

func newServer(addr string, cc *connectors) (s *server) {
	gs := grpc.NewServer()
	s = &server{addr, gs, cc}
	userpb.RegisterUserServiceServer(gs, s)
	healthpb.RegisterHealthServer(gs, s)
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

func (s *server) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

func (s *server) Watch(*healthpb.HealthCheckRequest, healthpb.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "Unimplemented")
}

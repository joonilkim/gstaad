package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	postpb "gstaad/src/postservice/pb"

	_ "google.golang.org/grpc/encoding/gzip" // Install the gzip compressor
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type server struct {
	log        *zap.SugaredLogger
	addr       string
	grpc       *grpc.Server
	connectors *connectors
}

func newServer(addr string, cc *connectors) (s *server) {
	// TODO: support SetLoggerV2
	grpclog.SetLogger(zapgrpc.NewLogger(logger, zapgrpc.WithDebug()))
	gs := grpc.NewServer()

	s = &server{logger.Sugar(), addr, gs, cc}
	postpb.RegisterPostServiceServer(gs, s)
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
		s.log.Infof("listening grpc %s", s.addr)
		er := s.grpc.Serve(cc)
		if er != nil && er != http.ErrServerClosed {
			s.log.Fatalf("Failed: %s\n", er)
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

package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	addr = "unix:///tmp/gstaad.user.test.sock"
)

func initTest(t *testing.T) error {
	return os.Remove(addr[7:])
}

func startMockGrpcServer(cc *connectors) *grpc.Server {
	gconn, er := net.Listen("unix", addr[7:])
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

func mockRouter() *http.ServeMux {
	return restServer(context.Background(), addr)
}

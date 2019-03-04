package main

import (
	"context"
	_log "log"
	"net"
	"net/http"
	"os"
	"testing"

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
		_log.Printf("listening grpc %s", addr)
		er := gs.Serve(gconn)
		if er != nil && er != http.ErrServerClosed {
			_log.Printf("Failed: %s\n", er)
		}
	}()
	return gs
}

func mockRouter() http.Handler {
	return restServer(context.Background(), addr)
}

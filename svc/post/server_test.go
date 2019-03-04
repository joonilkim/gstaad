package main

import (
	"net"
	"net/http"
	"os"

	"google.golang.org/grpc"
)

const mockSock = "unix:///tmp/test.sock"

func startMockServer() (s *server) {
	os.Remove(mockSock[7:])

	s = newServer(mockSock)
	cc, er := net.Listen("unix", mockSock[7:])
	if er != nil {
		panic(er)
	}

	go func() {
		er := s.Server.Serve(cc)
		if er != nil && er != http.ErrServerClosed {
			panic(er)
		}
	}()
	return
}

func mockConn() *grpc.ClientConn {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	cc, er := grpc.Dial(mockSock, opts...)
	if er != nil {
		panic(er)
	}
	return cc
}

package main

import (
	poststub "gstaad/pkg/proto/post"
	"os"

	"google.golang.org/grpc"
)

type connectors struct {
	post poststub.PostServiceClient
}

func NewConnectors() (*connectors, error) {
	post, er := newPostServiceClient()
	if er != nil {
		return nil, er
	}

	return &connectors{post}, nil
}

func newPostServiceClient() (poststub.PostServiceClient, error) {
	addr := os.Getenv("POST_ADDR")
	if addr == "" {
		addr = "post.gstaad:9001"
	}

	opts := []grpc.DialOption{grpc.WithInsecure()}
	cc, er := grpc.Dial(addr, opts...)
	if er != nil {
		return nil, er
	}
	return poststub.NewPostServiceClient(cc), nil
}

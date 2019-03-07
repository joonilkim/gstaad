package main

import (
	"google.golang.org/grpc"

	postpb "gstaad/src/postservice/pb"
)

type connectors struct {
	postservice postpb.PostServiceClient
}

func newConnectors() (*connectors, error) {
	ps, er := newPostServiceClient()
	if er != nil {
		return nil, er
	}

	return &connectors{ps}, nil
}

func newPostServiceClient() (postpb.PostServiceClient, error) {
	addr := mustGetenv("POSTSERVICE")

	opts := []grpc.DialOption{grpc.WithInsecure()}
	cc, er := grpc.Dial(addr, opts...)
	if er != nil {
		return nil, er
	}
	return postpb.NewPostServiceClient(cc), nil
}

package main

import (
	"context"
	pb "gstaad/pkg/proto/post"
)

type server struct{}

func (s *server) Create(c context.Context, in *pb.PostRequest) (*pb.PostReply, error) {
	return &pb.PostReply{Message: in.Name}, nil
}

package main

import (
	"context"
	pb "gstaad/pkg/proto/post"
	"time"

	empty "github.com/golang/protobuf/ptypes/empty"
)

type server struct{}

func (s *server) Create(c context.Context, in *pb.CreateRequest) (*pb.MutationReply, error) {
	return &pb.MutationReply{Result: true}, nil
}

func (s *server) All(c context.Context, in *empty.Empty) (*pb.PostsReply, error) {
	author := pb.User{
		Id:   "user1",
		Name: "user1",
	}

	posts := []*pb.Post{
		&pb.Post{
			Id:        "post1",
			Author:    &author,
			Content:   "post1",
			CreatedAt: time.Now().UnixNano(),
		},
		&pb.Post{
			Id:        "post2",
			Author:    &author,
			Content:   "post2",
			CreatedAt: time.Now().UnixNano(),
		},
	}
	return &pb.PostsReply{
		Items: posts,
	}, nil
}

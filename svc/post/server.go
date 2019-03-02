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
			CreatedAt: time.Now().UnixNano() / 1000,
		},
		&pb.Post{
			Id:        "post2",
			Author:    &author,
			Content:   "post2",
			CreatedAt: time.Now().UnixNano() / 1000,
		},
	}
	return &pb.PostsReply{
		Items: posts,
	}, nil
}

func (s *server) Count(c context.Context, in *empty.Empty) (*pb.CountReply, error) {
	return &pb.CountReply{
		Count: 10,
	}, nil
}

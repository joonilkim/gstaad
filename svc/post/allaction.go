package main

import (
	"context"
	"time"

	stub "gstaad/pkg/proto/post"

	"github.com/golang/protobuf/ptypes/empty"
)

func (s *server) All(c context.Context, in *empty.Empty) (*stub.PostsReply, error) {
	author := stub.User{
		Id:   "user1",
		Name: "user1",
	}

	posts := []*stub.Post{
		&stub.Post{
			Id:        "post1",
			Author:    &author,
			Content:   "post1",
			CreatedAt: time.Now().UnixNano() / 1000,
		},
		&stub.Post{
			Id:        "post2",
			Author:    &author,
			Content:   "post2",
			CreatedAt: time.Now().UnixNano() / 1000,
		},
	}
	return &stub.PostsReply{
		Items: posts,
	}, nil
}

package main

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"

	postpb "gstaad/src/postservice/pb"
)

func (s *server) All(c context.Context, in *empty.Empty) (*postpb.PostsReply, error) {
	author := postpb.User{
		Id:   "user1",
		Name: "user1",
	}

	posts := []*postpb.Post{
		&postpb.Post{
			Id:        "post1",
			Author:    &author,
			Content:   "post1",
			CreatedAt: time.Now().UnixNano() / 1000,
		},
		&postpb.Post{
			Id:        "post2",
			Author:    &author,
			Content:   "post2",
			CreatedAt: time.Now().UnixNano() / 1000,
		},
	}
	return &postpb.PostsReply{
		Items: posts,
	}, nil
}

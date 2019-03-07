package main

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	postpb "gstaad/src/postservice/pb"
)

func (s *server) Count(c context.Context, in *empty.Empty) (*postpb.CountReply, error) {
	return &postpb.CountReply{
		Count: 10,
	}, nil
}

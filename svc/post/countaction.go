package main

import (
	"context"

	stub "gstaad/pkg/proto/post"

	"github.com/golang/protobuf/ptypes/empty"
)

func (s *server) Count(c context.Context, in *empty.Empty) (*stub.CountReply, error) {
	return &stub.CountReply{
		Count: 10,
	}, nil
}

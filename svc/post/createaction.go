package main

import (
	"context"

	stub "gstaad/pkg/proto/post"
)

func (s *server) Create(c context.Context, in *stub.CreateRequest) (*stub.MutationReply, error) {
	return &stub.MutationReply{Result: true}, nil
}

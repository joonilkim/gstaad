package main

import (
	"context"

	postpb "gstaad/src/postservice/pb"
)

func (s *server) Create(c context.Context, in *postpb.CreateRequest) (*postpb.MutationReply, error) {
	return &postpb.MutationReply{Result: true}, nil
}

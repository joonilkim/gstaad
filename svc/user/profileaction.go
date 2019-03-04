package main

import (
	"context"

	stub "gstaad/pkg/proto/user"

	empty "github.com/golang/protobuf/ptypes/empty"
)

func (s *server) Profile(c context.Context, in *empty.Empty) (*stub.ProfileReply, error) {
	var count int32
	r, er := s.connectors.post.Count(c, &empty.Empty{})
	if er == nil {
		count = r.Count
	}

	return &stub.ProfileReply{
		User: &stub.User{
			Id:   "testuser",
			Name: "testname",
		},
		PostCount: count,
	}, nil
}

package main

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"

	userpb "gstaad/src/userservice/pb"

	empty "github.com/golang/protobuf/ptypes/empty"
)

func (s *server) Login(c context.Context, in *userpb.LoginRequest) (*userpb.LoginReply, error) {
	expiresIn, _ := ptypes.TimestampProto(time.Now().AddDate(0, 0, 1))
	return &userpb.LoginReply{
		Token: &userpb.Token{
			AccessToken: "testtoken",
			ExpiresIn:   expiresIn,
		},
	}, nil
}

func (s *server) Profile(c context.Context, in *empty.Empty) (*userpb.ProfileReply, error) {
	var count int32
	r, er := s.connectors.postservice.Count(c, &empty.Empty{})
	if er == nil {
		count = r.Count
	}

	return &userpb.ProfileReply{
		User: &userpb.User{
			Id:   "testuser",
			Name: "testname",
		},
		PostCount: count,
	}, nil
}

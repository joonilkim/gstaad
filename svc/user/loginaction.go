package main

import (
	"context"
	"time"

	stub "gstaad/pkg/proto/user"

	"github.com/golang/protobuf/ptypes"
)

func (s *server) Login(c context.Context, in *stub.LoginRequest) (*stub.LoginReply, error) {
	expiresIn, _ := ptypes.TimestampProto(time.Now().AddDate(0, 0, 1))
	return &stub.LoginReply{
		Token: &stub.Token{
			AccessToken: "testtoken",
			ExpiresIn:   expiresIn,
		},
	}, nil
}

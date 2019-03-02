package main

import (
	"context"
	pb "gstaad/pkg/proto/user"
	"time"

	"github.com/golang/protobuf/ptypes"

	empty "github.com/golang/protobuf/ptypes/empty"
)

type server struct {
	connectors *connectors
}

func (s *server) Login(c context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	expiresIn, _ := ptypes.TimestampProto(time.Now().AddDate(0, 0, 1))
	return &pb.LoginReply{
		Token: &pb.Token{
			AccessToken: "testtoken",
			ExpiresIn:   expiresIn,
		},
	}, nil
}

func (s *server) Profile(c context.Context, in *empty.Empty) (*pb.ProfileReply, error) {
	var count int32
	r, er := s.connectors.post.Count(c, &empty.Empty{})
	if er == nil {
		count = r.Count
	}

	return &pb.ProfileReply{
		User: &pb.User{
			Id:   "testuser",
			Name: "testname",
		},
		PostCount: count,
	}, nil
}

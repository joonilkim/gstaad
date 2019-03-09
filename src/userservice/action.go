package main

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"

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
	cnt := s.postCount(c)

	return &userpb.ProfileReply{
		User: &userpb.User{
			Id:   "testuser",
			Name: "testname",
		},
		PostCount: cnt,
	}, nil
}

func (s *server) postCount(c context.Context) (cnt int32) {
	c, cancel := context.WithTimeout(c, time.Millisecond*100)
	defer cancel()

	// just a sample to show how to use compressor
	// actually no need to compress for this
	r, er := s.cc.postSvc.Count(c, &empty.Empty{}, grpc.UseCompressor(gzip.Name))
	if er != nil {
		logger.Errorf("Fail to connect postservice: %v", er)
	} else {
		cnt = r.Count
	}
	return
}

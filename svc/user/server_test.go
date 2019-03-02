package main

import (
	"context"
	"testing"
	"time"

	poststub "gstaad/pkg/proto/post"
	postmock "gstaad/pkg/proto/post/mock"
	pb "gstaad/pkg/proto/user"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
	empty "github.com/golang/protobuf/ptypes/empty"
	assert "github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	post := postmock.NewMockPostServiceClient(ctrl)
	cc := &connectors{post}

	initTest(t)
	gs := startMockGrpcServer(cc)
	defer gs.Stop()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, er := grpc.Dial(addr, opts...)
	assert.NoError(t, er)
	defer conn.Close()

	user := pb.NewUserServiceClient(conn)

	t.Run("login", func(t *testing.T) {
		name := "testname"
		r, er := user.Login(context.Background(), &pb.LoginRequest{Name: name})
		assert.NoError(t, er)
		assert.NotEmpty(t, true, r.Token.AccessToken)
		expiresIn, _ := ptypes.Timestamp(r.Token.ExpiresIn)
		assert.True(t, expiresIn.After(time.Now()))
	})

	t.Run("profile", func(t *testing.T) {
		post.EXPECT().Count(gomock.Any(), &empty.Empty{}).
			Return(&poststub.CountReply{Count: 10}, nil)

		r, er := user.Profile(context.Background(), &empty.Empty{})
		assert.NoError(t, er)
		assert.NotEmpty(t, r.User.Id)
		assert.NotEmpty(t, r.User.Name)
		assert.NotEmpty(t, r.PostCount)
	})

}

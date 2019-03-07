package main

import (
	"context"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	postpb "gstaad/src/postservice/pb"
	userpb "gstaad/src/userservice/pb"

	assert "github.com/stretchr/testify/require"
)

type mockPostServiceClient struct{}

func (*mockPostServiceClient) Create(ctx context.Context, in *postpb.CreateRequest, opts ...grpc.CallOption) (*postpb.MutationReply, error) {
	return nil, nil
}

func (*mockPostServiceClient) All(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*postpb.PostsReply, error) {
	return nil, nil
}

func (*mockPostServiceClient) Count(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*postpb.CountReply, error) {
	return &postpb.CountReply{Count: 10}, nil
}

func TestAction(t *testing.T) {
	postcli := &mockPostServiceClient{}

	gs := startMockServer(&connectors{postcli})
	defer gs.grpc.Stop()

	cc := mockConn()
	defer cc.Close()
	user := userpb.NewUserServiceClient(cc)

	t.Run("login action", func(t *testing.T) {
		t.Run("should pass", func(t *testing.T) {
			name := "testname"
			r, er := user.Login(context.Background(), &userpb.LoginRequest{Name: name})
			assert.NoError(t, er)
			assert.NotEmpty(t, true, r.Token.AccessToken)
			expiresIn, _ := ptypes.Timestamp(r.Token.ExpiresIn)
			assert.True(t, expiresIn.After(time.Now()))
		})
	})

	t.Run("login action", func(t *testing.T) {
		t.Run("should pass", func(t *testing.T) {
			r, er := user.Profile(context.Background(), &empty.Empty{})
			assert.NoError(t, er)
			assert.NotNil(t, r)
			assert.NotNil(t, r.User)

			assert.NotEmpty(t, r.User.Id)
			assert.NotEmpty(t, r.User.Name)
			assert.NotEmpty(t, r.PostCount)
		})
	})
}

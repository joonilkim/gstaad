package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"gstaad/src/userservice/cognito"
	"gstaad/src/userservice/types"
	"testing"
	"time"

	"google.golang.org/grpc/metadata"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

type mockAuthSvc struct {
	email  string
	pw     string
	expect *cognito.TokenPayload
}

func (s *mockAuthSvc) Signup(email, pw string) error {
	if s.email == email {
		return status.Error(codes.AlreadyExists, "AlreadyExists")
	}

	s.email = email
	s.pw = pw
	return nil
}

func (s *mockAuthSvc) Login(tk *cognito.TokenPayload, email, pw string) error {
	if s.email != email {
		return status.Error(codes.NotFound, "NotFound")
	}

	tk.AccessToken = s.expect.AccessToken
	tk.ExpiresIn = s.expect.ExpiresIn
	return nil
}

func (c *mockAuthSvc) Unregister(token string) error {
	return nil
}

func TestAction(t *testing.T) {
	authSvc := &mockAuthSvc{
		expect: &cognito.TokenPayload{
			AccessToken: "1234",
			ExpiresIn:   time.Now().Add(10 * time.Minute).Unix(),
		},
	}
	postSvc := &mockPostServiceClient{}

	gs := startMockServer(&connectors{authSvc, postSvc})
	defer gs.grpc.Stop()

	cc := mockConn()
	defer cc.Close()
	user := userpb.NewUserServiceClient(cc)

	email, pw := "aa@aa.com", "asdf123$"
	jp := types.JWTPayload{Email: email}

	t.Run("login action", func(t *testing.T) {
		t.Run("should signup", func(t *testing.T) {
			p := userpb.SignupRequest{Email: email, Password: pw}
			r, er := user.Signup(context.Background(), &p)
			assert.NoError(t, er)
			assert.Equal(t, true, r.Result)
		})

		t.Run("should login", func(t *testing.T) {
			p := userpb.LoginRequest{Email: email, Password: pw}
			r, er := user.Login(context.Background(), &p)

			assert.NoError(t, er)
			assert.NotEmpty(t, true, r.Token.AccessToken)
			assert.NotEmpty(t, r.Token.ExpiresIn)
		})
	})

	t.Run("profile action", func(t *testing.T) {
		t.Run("should pass", func(t *testing.T) {
			ctx, er := withJWTPayload(context.Background(), jp)
			assert.NoError(t, er)

			d, er := user.Profile(ctx, &empty.Empty{})
			assert.NoError(t, er)
			assert.NotNil(t, d)
			assert.NotNil(t, d.User)

			assert.Equal(t, email, d.User.Email)
			assert.NotEmpty(t, d.PostCount)
		})
	})
}

func withJWTPayload(ctx context.Context, p types.JWTPayload) (context.Context, error) {
	b := &bytes.Buffer{}
	if er := encodePayload(b, p); er != nil {
		return nil, er
	}

	md := metadata.Pairs("x-jwt-payload", string(b.Bytes()))
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx, nil
}

func encodePayload(b *bytes.Buffer, p types.JWTPayload) error {
	enc := base64.NewEncoder(base64.StdEncoding, b)
	defer enc.Close()
	return json.NewEncoder(enc).Encode(p)
}

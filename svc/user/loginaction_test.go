package main

import (
	"context"
	"testing"
	"time"

	pb "gstaad/pkg/proto/user"

	"github.com/golang/protobuf/ptypes"
	assert "github.com/stretchr/testify/require"
)

func TestLoginAction(t *testing.T) {
	gs := startMockServer(&connectors{})
	defer gs.grpc.Stop()

	cc := mockConn()
	defer cc.Close()
	user := pb.NewUserServiceClient(cc)

	t.Run("should pass", func(t *testing.T) {
		name := "testname"
		r, er := user.Login(context.Background(), &pb.LoginRequest{Name: name})
		assert.NoError(t, er)
		assert.NotEmpty(t, true, r.Token.AccessToken)
		expiresIn, _ := ptypes.Timestamp(r.Token.ExpiresIn)
		assert.True(t, expiresIn.After(time.Now()))
	})
}

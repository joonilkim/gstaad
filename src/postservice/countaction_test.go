package main

import (
	"context"
	"testing"

	postpb "gstaad/src/postservice/pb"

	empty "github.com/golang/protobuf/ptypes/empty"
	assert "github.com/stretchr/testify/require"
)

func TestCountAction(t *testing.T) {
	gs := startMockServer()
	defer gs.grpc.Stop()

	cc := mockConn()
	defer cc.Close()
	post := postpb.NewPostServiceClient(cc)

	t.Run("should pass", func(t *testing.T) {
		r, er := post.Count(context.Background(), &empty.Empty{})
		assert.NoError(t, er)
		assert.NotEmpty(t, r.Count)
	})
}
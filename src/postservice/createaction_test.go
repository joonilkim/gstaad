package main

import (
	"context"
	"testing"

	postpb "gstaad/src/postservice/pb"

	assert "github.com/stretchr/testify/require"
)

func TestCreateAction(t *testing.T) {
	gs := startMockServer()
	defer gs.grpc.Stop()

	cc := mockConn()
	defer cc.Close()
	post := postpb.NewPostServiceClient(cc)

	t.Run("should pass", func(t *testing.T) {
		content := "testname"
		r, er := post.Create(context.Background(), &postpb.CreateRequest{Content: content})
		assert.NoError(t, er)
		assert.Equal(t, true, r.Result)
	})
}

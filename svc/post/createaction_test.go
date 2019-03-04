package main

import (
	"context"
	"testing"

	pb "gstaad/pkg/proto/post"

	assert "github.com/stretchr/testify/require"
)

func TestCreateAction(t *testing.T) {
	gs := startMockServer()
	defer gs.Server.Stop()

	cc := mockConn()
	defer cc.Close()
	post := pb.NewPostServiceClient(cc)

	t.Run("should pass", func(t *testing.T) {
		content := "testname"
		r, er := post.Create(context.Background(), &pb.CreateRequest{Content: content})
		assert.NoError(t, er)
		assert.Equal(t, true, r.Result)
	})
}

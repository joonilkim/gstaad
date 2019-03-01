package main

import (
	"context"
	"testing"

	pb "gstaad/pkg/proto/post"

	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestPost(t *testing.T) {
	addr := "unix:///tmp/test.sock"
	gsv := startGrpc(addr)
	defer gsv.Stop()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, er := grpc.Dial(addr, opts...)
	assert.NoError(t, er)
	defer conn.Close()

	post := pb.NewPostServiceClient(conn)

	t.Run("create", func(t *testing.T) {
		content := "testname"
		r, er := post.Create(context.Background(), &pb.CreateRequest{Content: content})
		assert.NoError(t, er)
		assert.Equal(t, true, r.Result)
	})

	t.Run("all", func(t *testing.T) {
		r, er := post.All(context.Background(), &empty.Empty{})
		assert.NoError(t, er)
		assert.NotEmpty(t, r.Items)
	})

}

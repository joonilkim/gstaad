package main

import (
	"context"
	"testing"

	pb "gstaad/pkg/proto/post"

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

	post := pb.NewPostClient(conn)

	t.Run("create", func(t *testing.T) {
		name := "testname"
		r, er := post.Create(context.Background(), &pb.PostRequest{Name: name})
		assert.NoError(t, er)
		assert.Equal(t, name, r.Message)
	})

}

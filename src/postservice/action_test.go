package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"gstaad/src/postservice/types"
	"testing"

	postpb "gstaad/src/postservice/pb"

	empty "github.com/golang/protobuf/ptypes/empty"
	assert "github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func TestAction(t *testing.T) {

	t.Run("post", func(t *testing.T) {
		gs := startMockServer()
		defer gs.grpc.Stop()

		cc := mockConn()
		defer cc.Close()
		post := postpb.NewPostServiceClient(cc)

		jp := types.JWTPayload{Email: "aa@aa.com"}

		t.Run("should create", func(t *testing.T) {
			ctx, er := withJWTPayload(context.Background(), jp)
			assert.NoError(t, er)

			content := "testname"
			r, er := post.Create(ctx, &postpb.CreateRequest{Content: content})
			assert.NoError(t, er)
			assert.Equal(t, true, r.Result)
		})

		t.Run("should return all", func(t *testing.T) {
			ctx, er := withJWTPayload(context.Background(), jp)
			assert.NoError(t, er)

			r, er := post.All(ctx, &empty.Empty{})
			assert.NoError(t, er)
			assert.NotEmpty(t, r.Items)
		})

		t.Run("should return count", func(t *testing.T) {
			ctx, er := withJWTPayload(context.Background(), jp)
			assert.NoError(t, er)

			r, er := post.Count(ctx, &empty.Empty{})
			assert.NoError(t, er)
			assert.NotEmpty(t, r.Count)
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

package main

import (
	"context"
	"testing"

	poststub "gstaad/pkg/proto/post"
	postmock "gstaad/pkg/proto/post/mock"
	pb "gstaad/pkg/proto/user"

	"github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	assert "github.com/stretchr/testify/require"
)

func TestProfileAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	postsvc := postmock.NewMockPostServiceClient(ctrl)

	gs := startMockServer(&connectors{postsvc})
	defer gs.Server.Stop()

	cc := mockConn()
	defer cc.Close()
	user := pb.NewUserServiceClient(cc)

	t.Run("profile", func(t *testing.T) {
		postsvc.EXPECT().Count(gomock.Any(), &empty.Empty{}).
			Return(&poststub.CountReply{Count: 10}, nil)

		r, er := user.Profile(context.Background(), &empty.Empty{})
		assert.NoError(t, er)
		assert.NotNil(t, r)
		assert.NotNil(t, r.User)

		assert.NotEmpty(t, r.User.Id)
		assert.NotEmpty(t, r.User.Name)
		assert.NotEmpty(t, r.PostCount)
	})
}

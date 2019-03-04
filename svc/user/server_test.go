package main

import (
	"net"
	"net/http"
	"os"

	"google.golang.org/grpc"
)

/*
func TestUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	post := postmock.NewMockPostServiceClient(ctrl)
	cc := &connectors{post}

	initTest(t)
	gs := startMockGrpcServer(cc)
	defer gs.Stop()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, er := grpc.Dial(addr, opts...)
	assert.NoError(t, er)
	defer conn.Close()

	user := pb.NewUserServiceClient(conn)

	t.Run("login", func(t *testing.T) {
		name := "testname"
		r, er := user.Login(context.Background(), &pb.LoginRequest{Name: name})
		assert.NoError(t, er)
		assert.NotEmpty(t, true, r.Token.AccessToken)
		expiresIn, _ := ptypes.Timestamp(r.Token.ExpiresIn)
		assert.True(t, expiresIn.After(time.Now()))
	})

	t.Run("profile", func(t *testing.T) {
		post.EXPECT().Count(gomock.Any(), &empty.Empty{}).
			Return(&poststub.CountReply{Count: 10}, nil)

		r, er := user.Profile(context.Background(), &empty.Empty{})
		assert.NoError(t, er)
		assert.NotEmpty(t, r.User.Id)
		assert.NotEmpty(t, r.User.Name)
		assert.NotEmpty(t, r.PostCount)
	})

}
*/

const mockSock = "unix:///tmp/test.sock"

func startMockServer(ccs *connectors) (s *server) {
	os.Remove(mockSock[7:])

	s = newServer(mockSock, ccs)
	cc, er := net.Listen("unix", mockSock[7:])
	if er != nil {
		panic(er)
	}

	go func() {
		er := s.Server.Serve(cc)
		if er != nil && er != http.ErrServerClosed {
			panic(er)
		}
	}()
	return
}

func mockConn() *grpc.ClientConn {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	cc, er := grpc.Dial(mockSock, opts...)
	if er != nil {
		panic(er)
	}
	return cc
}

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"

	"gstaad/src/userservice/cognito"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"

	postpb "gstaad/src/postservice/pb"
)

type connectors struct {
	authSvc cognito.CognitoClient
	postSvc postpb.PostServiceClient
}

func newConnectors(ctx context.Context) (cc *connectors) {
	sess := session.Must(session.NewSession())

	if os.Getenv("APP_ENV") == "production" {
		resolver.SetDefaultScheme("dns")
	}

	// connections can be shared across multiple clients via multiplexing
	postcc := new(grpc.ClientConn)
	mustGetConn(ctx, &postcc, mustGetenv("POSTSERVICE"))

	cc = &connectors{}
	cc.authSvc = cognito.NewCognitoClient(sess)
	cc.postSvc = postpb.NewPostServiceClient(postcc)
	return
}

func mustGetConn(ctx context.Context, cc **grpc.ClientConn, addr string) {
	var er error
	*cc, er = grpc.DialContext(ctx, addr, defaultDialOpts()...)
	if er != nil {
		panic(fmt.Sprintf("grpc: failed to connect %s", addr))
	}
}

func defaultDialOpts() []grpc.DialOption {
	// use non-blocking
	// Dial returns a client that does not have any servers connected and continues
	// to watch for connection in the background. Failure occurs at RPC time
	// if no server has been found.
	return []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
		// use default backoff. max 2min
		// use default keepalive. inifinity and 20s ping intervals
	}
}

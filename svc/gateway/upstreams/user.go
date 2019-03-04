package upstreams

import (
	"context"

	stub "gstaad/pkg/proto/user"
	"gstaad/pkg/utils"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

const userServiceEndpoint = "user.gstaad:9000"

type userService struct{}

func (p *userService) apply(c context.Context, mux *runtime.ServeMux) error {
	u := utils.Getenv("USERSVC_ENDPOINT", userServiceEndpoint)

	op := []grpc.DialOption{grpc.WithInsecure()}
	return stub.RegisterUserServiceHandlerFromEndpoint(c, mux, u, op)
}

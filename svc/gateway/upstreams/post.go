package upstreams

import (
	"context"

	stub "gstaad/pkg/proto/post"
	"gstaad/pkg/utils"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

const postServiceEndpoint = "post.gstaad:9000"

type postService struct{}

func (p *postService) apply(c context.Context, mux *runtime.ServeMux) error {
	u := utils.Getenv("POSTSVC_ENDPOINT", postServiceEndpoint)

	op := []grpc.DialOption{grpc.WithInsecure()}
	return stub.RegisterPostServiceHandlerFromEndpoint(c, mux, u, op)
}

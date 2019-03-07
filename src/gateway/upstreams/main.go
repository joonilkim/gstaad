package upstreams

import (
	"context"
	"fmt"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type service interface {
	apply(c context.Context, mux *runtime.ServeMux) error
}

var services = []service{
	&postService{},
	&userService{},
}

func Apply(c context.Context, mux *runtime.ServeMux) error {
	es := make([]error, len(services))

	for i, s := range services {
		es[i] = s.apply(c, mux)
	}
	return joinErrors(es)
}

func joinErrors(es []error) error {
	ss := strings.Builder{}

	for _, e := range es {
		if e != nil {
			ss.WriteString(e.Error())
		}
	}
	return fmt.Errorf(ss.String())
}

module gstaad/src/userservice

require (
	github.com/golang/protobuf v1.3.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/pkg/errors v0.8.1 // indirect
	github.com/sirupsen/logrus v1.3.0
	github.com/stretchr/testify v1.3.0
	go.uber.org/atomic v1.3.2 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.9.1
	google.golang.org/genproto v0.0.0-20190306222511-6e86cb5d2f12
	google.golang.org/grpc v1.19.0
	gstaad/src/postservice v0.0.0
)

replace gstaad/src/postservice => github.com/joonilkim/gstaad/src/postservice v0.0.0

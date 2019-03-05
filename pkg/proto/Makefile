gen-mock:
	mkdir -p mock && \
		cd ${root} && \
		mockgen gstaad/pkg/proto/${service} ${mockClient} > \
		pkg/proto/${service}/mock/${service}.mock.go

gen-stub:
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		${service}.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		${service}.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:. \
		${service}.proto
	protoc -I/usr/local/include -I. \
		--include_imports --include_source_info \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--descriptor_set_out=${service}.pb \
		${service}.proto
	go generate .

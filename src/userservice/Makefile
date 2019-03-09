include .env
export

pbdir  := ../../pb
protos := userservice
googleapis := ${GOPATH}/src/googleapis

.PHONY: pb
pb: $(protos)

${googleapis}:
	@mkdir -p pb
	git clone --depth=1 https://github.com/googleapis/googleapis $@ > /dev/null

.PHONY: $(protos)
$(protos): ${googleapis}
	protoc --descriptor_set_out=pb/$@.pb \
		--include_imports --include_source_info \
		-I${pbdir} -I${GOPATH}/src/googleapis \
		${pbdir}/$@.proto
	protoc --go_out=plugins=grpc:pb \
		-I${pbdir} -I${GOPATH}/src/googleapis \
		${pbdir}/$@.proto

test:
	PORT=-1 go test

run:
	PORT=9100 POSTSERVICE=:9200 go run . -v=INFO

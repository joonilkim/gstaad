FROM golang:1.11.5

ENV GO111MODULE=on
ARG SERVICE_NAME=postservice
ARG SERVICE_PATH=/go/src/gstaad/src/${SERVICE_NAME}

RUN mkdir -p ${SERVICE_PATH} /go/bin

# caching layer
COPY go.mod go.sum ${SERVICE_PATH}/
WORKDIR ${SERVICE_PATH}
RUN go mod download

COPY . /go/src/gstaad/
RUN find /go/src/gstaad
RUN pwd
RUN go build -o /go/bin/gstaad .

WORKDIR /go/bin
RUN rm -rf /go/src/gstaad

RUN apk add --no-cache ca-certificates
RUN GRPC_HEALTH_PROBE_VERSION=v0.2.0 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

ENTRYPOINT ["/go/bin/gstaad"]

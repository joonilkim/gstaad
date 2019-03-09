FROM golang:1.11.5

ENV GO111MODULE=on
ARG SERVICE_NAME=userservice
ARG SERVICE_PATH=/go/src/gstaad/src/${SERVICE_NAME}

RUN GRPC_HEALTH_PROBE_VERSION=v0.2.1 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

RUN mkdir -p ${SERVICE_PATH} /go/bin

# caching layer
COPY go.mod go.sum ${SERVICE_PATH}/
WORKDIR ${SERVICE_PATH}
RUN go mod download

COPY . ${SERVICE_PATH}
RUN go build -o /go/bin/gstaad .

WORKDIR /go/bin
RUN rm -rf /go/src /go/pkg

CMD ["/go/bin/gstaad"]

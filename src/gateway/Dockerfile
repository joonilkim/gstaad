FROM golang:1.11.5

ENV GO111MODULE=on

RUN mkdir -p /go/src/gstaad /go/bin

# caching layer
COPY go.mod go.sum /go/src/gstaad/
WORKDIR /go/src/gstaad
RUN go mod download

COPY . /go/src/gstaad/
WORKDIR /go/src/gstaad/svc/gateway
RUN go build -o /go/bin/gstaad .

WORKDIR /go/bin
RUN rm -rf /go/src/gstaad
EXPOSE 8080

ENTRYPOINT ["/go/bin/gstaad"]

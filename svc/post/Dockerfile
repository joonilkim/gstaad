FROM golang:1.11.5

ENV GO111MODULE=on

RUN mkdir -p /go/src/gstaad /go/bin
COPY . /go/src/gstaad/

WORKDIR /go/src/gstaad
RUN go build -o /go/bin/gstaad svc/post/main.go

WORKDIR /go/bin
RUN rm -rf /go/src/gstaad
EXPOSE 9000

ENTRYPOINT ["/go/bin/gstaad"]

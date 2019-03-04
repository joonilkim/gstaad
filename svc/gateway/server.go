package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	mw "gstaad/svc/gateway/middlewares"
	"gstaad/svc/gateway/upstreams"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type server struct {
	log     *log.Logger
	context context.Context
	server  *http.Server
}

func newServer(addr string, l *log.Logger) (sv *server) {
	c := context.Background()

	r := newRouter(c)
	r = withLogging(l, r)
	r = withH2C(r)
	s := &http.Server{
		Addr:           addr,
		Handler:        r,
		ErrorLog:       l,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &server{l, c, s}
}

func (s *server) run() {
	go func() {
		logger.Infof("listening: %s", s.server.Addr)
		er := s.server.ListenAndServe()
		if er != nil && er != http.ErrServerClosed {
			logger.Fatalf("Failed: %s\n", er)
		}
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	logger.Info("Stopping operation...")

	ctx, cancel := context.WithTimeout(s.context, 5*time.Second)
	defer cancel()

	s.server.SetKeepAlivesEnabled(false)
	if er := s.server.Shutdown(ctx); er != nil {
		logger.Fatalf("Shutdown down: %v", er)
	}
	<-ctx.Done()
}

func grpcRoutes(c context.Context) http.Handler {
	r := runtime.NewServeMux()
	upstreams.Apply(c, r)
	return r
}

func newRouter(c context.Context) http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/ping", servePing)
	r.Handle("/", grpcRoutes(c))
	return r
}

func withLogging(l *log.Logger, h http.Handler) http.Handler {
	nextID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return mw.Tracing(nextID)(mw.Logging(l)(h))
}

func withH2C(h http.Handler) http.Handler {
	return h2c.NewHandler(h, &http2.Server{})
}

func servePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

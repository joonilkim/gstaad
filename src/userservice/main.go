package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"go.uber.org/zap"
)

const (
	Version = "0.0.0+3"
)

var (
	port     = mustGetenv("PORT")
	logger   *zap.Logger
	loglevel = zap.LevelFlag("v", zap.InfoLevel, "")
)

func init() {
	var cfg zap.Config

	if os.Getenv("APP_ENV") == "production" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
		//cfg.Level.SetLevel(zapcore.DebugLevel)
	}

	flag.Parse()
	cfg.Level.SetLevel(*loglevel)

	var er error
	logger, er = cfg.Build()
	must(er)
}

func main() {
	defer logger.Sync()

	ctx := context.Background()
	cc := newConnectors(ctx)

	addr := ":" + port
	s := newServer(addr, cc)
	s.run()
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", k))
	}
	return v
}

func must(e error) {
	if e != nil {
		panic(e)
	}
}

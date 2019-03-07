package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	Version = "0.0.0+1"
)

var (
	port   = mustGetenv("PORT")
	logger = logrus.New()
)

func init() {
	if os.Getenv("APP_ENV") == "production" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{})
	}
}

func main() {
	cc, er := newConnectors()
	must(er)

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

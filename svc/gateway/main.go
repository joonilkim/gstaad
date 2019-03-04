package main

import (
	"log"
	"os"

	"gstaad/pkg/utils"

	"github.com/sirupsen/logrus"
)

const (
	Version = "0.0.0+1"
)

var (
	port   = utils.Getenv("PORT", "8080")
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
	logw := logger.Writer()
	defer logw.Close()

	addr := ":" + port
	s := newServer(addr, log.New(logw, "", 0))
	s.run()
}

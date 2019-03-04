package main

import (
	"os"

	"gstaad/pkg/utils"

	"github.com/sirupsen/logrus"
)

const (
	Version = "0.0.0+1"
)

var (
	port   = utils.Getenv("PORT", "9000")
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
	addr := ":" + port
	s := newServer(addr)
	s.run()
}

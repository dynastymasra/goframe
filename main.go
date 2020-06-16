package main

import (
	"github.com/dynastymasra/goframe/config"
	"github.com/sirupsen/logrus"
)

func init() {
	config.Load()
	config.Logger().Setup()
}

func main() {
	logrus.Info("test")
}

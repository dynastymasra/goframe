package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
	"gopkg.in/tylerb/graceful.v1"
)

func Run(server *graceful.Server, port string, router *RouterInstance) error {
	log := logrus.WithFields(logrus.Fields{
		"port":        port,
		"serviceName": router.ServiceName,
	})

	log.Infoln("Start run web application")

	muxRouter := router.Router()

	server.Server = &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: handlers.RecoveryHandler(
			handlers.PrintRecoveryStack(true),
			handlers.RecoveryLogger(logrus.StandardLogger()),
		)(muxRouter),
	}

	if err := server.ListenAndServe(); err != nil {
		log.WithError(err).Errorln("Failed run web application")
		return err
	}

	log.Infoln("Web application is running")

	return nil
}

package controller

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/goframe/config"

	"github.com/dynastymasra/cookbook"
	"github.com/dynastymasra/cookbook/message"
	"github.com/sirupsen/logrus"
)

// Remove unused params
func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(cookbook.HContentType, cookbook.HJSONTypeUTF8)
		w.Header().Set(cookbook.XRequestID, fmt.Sprintf("%v", r.Context().Value(config.RequestID)))

		log := logrus.WithFields(logrus.Fields{
			config.RequestID:    r.Context().Value(config.RequestID),
			config.JServiceName: config.ServiceName,
			config.JVersion:     config.Version,
		})

		if err := config.Neo4J().Ping(); err != nil {
			log.WithError(err).Errorln("Failed connect to Neo4J database")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), message.ErrDatabaseUnavailableCode).Stringify())
			return
		}

		if err := config.Postgres().Ping(); err != nil {
			log.WithError(err).Errorln("Failed connect to Postgres database")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), message.ErrDatabaseUnavailableCode).Stringify())
			return
		}

		if err := config.Elasticsearch().Ping(); err != nil {
			log.WithError(err).Errorln("Elasticsearch database connection has an error")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), message.ErrDatabaseUnavailableCode).Stringify())
			return
		}

		if err := config.Redis().Ping(); err != nil {
			log.WithError(err).Errorln("Failed connect to redis database")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), message.ErrDatabaseUnavailableCode).Stringify())
			return
		}

		if err := config.MongoDB().Ping(); err != nil {
			log.WithError(err).Errorln("Failed connect to mongo database")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), message.ErrDatabaseUnavailableCode).Stringify())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, cookbook.SuccessResponse().Stringify())
	}
}

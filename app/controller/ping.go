package controller

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/goframe/config"

	"github.com/elastic/go-elasticsearch/v7"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dynastymasra/cookbook"
	"github.com/sirupsen/logrus"
)

// Remove unused params
func Ping(db *gorm.DB, client *mongo.Client, esClient *elasticsearch.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(cookbook.HContentType, cookbook.HJSONTypeUTF8)
		w.Header().Set(cookbook.XRequestID, fmt.Sprintf("%v", r.Context().Value(config.RequestID)))

		log := logrus.WithFields(logrus.Fields{
			config.RequestID:    r.Context().Value(config.RequestID),
			config.JServiceName: config.ServiceName,
			config.JVersion:     config.Version,
		})

		dbClient, err := db.DB()
		if err != nil {
			log.WithError(err).Errorln("Failed get database")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), cookbook.ErrDatabaseUnavailableCode).Stringify())
			return
		}

		if err := dbClient.Ping(); err != nil {
			log.WithError(err).Errorln("Failed ping database")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), cookbook.ErrDatabaseUnavailableCode).Stringify())
			return
		}

		if err := client.Ping(r.Context(), nil); err != nil {
			log.WithError(err).Errorln("Failed ping mongo database")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), cookbook.ErrDatabaseUnavailableCode).Stringify())
			return
		}

		res, err := esClient.Ping()
		if err != nil {
			log.WithError(err).Errorln("Failed ping elasticsearch database")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), cookbook.ErrDatabaseUnavailableCode).Stringify())
			return
		}

		if res.IsError() {
			log.WithError(err).Errorln("Failed ping elasticsearch is error")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(res.String(), cookbook.ErrDatabaseUnavailableCode).Stringify())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, cookbook.SuccessResponse().Stringify())
	}
}

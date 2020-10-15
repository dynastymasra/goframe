package controller

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dynastymasra/goframe/config"

	"github.com/dynastymasra/cookbook"
	"github.com/sirupsen/logrus"
)

// Remove unused params
func Ping(db *gorm.DB, client *mongo.Client, esClient *elasticsearch.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		dbClient, err := db.DB()
		if err != nil {
			logrus.WithError(err).Errorln("Failed get database")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), r.Context().Value(config.RequestID)).Stringify())
			return
		}

		if err := dbClient.Ping(); err != nil {
			logrus.WithError(err).Errorln("Failed ping database")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), r.Context().Value(config.RequestID)).Stringify())
			return
		}

		if err := client.Ping(r.Context(), nil); err != nil {
			logrus.WithError(err).Errorln("Failed ping mongo database")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), r.Context().Value(config.RequestID)).Stringify())
			return
		}

		res, err := esClient.Ping()
		if err != nil {
			logrus.WithError(err).Errorln("Failed ping elasticsearch database")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), r.Context().Value(config.RequestID)).Stringify())
			return
		}

		if res.IsError() {
			logrus.WithError(err).Errorln("Failed ping elasticsearch is error")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(res.String(), r.Context().Value(config.RequestID)).Stringify())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, cookbook.SuccessResponse().Stringify())
	}
}

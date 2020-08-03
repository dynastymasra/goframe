package controller

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dynastymasra/goframe/config"

	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/cookbook"
	"github.com/sirupsen/logrus"
)

// Remove unused params
func Ping(db *gorm.DB, client *mongo.Client, esClient *elasticsearch.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := db.DB().Ping(); err != nil {
			logrus.WithError(err).Errorln("Failed ping database")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), r.Context().Value(config.RequestID)).Stringify())
			return
		}

		if err := client.Ping(r.Context(), nil); err != nil {
			logrus.WithError(err).Errorln("Failed ping database")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), r.Context().Value(config.RequestID)).Stringify())
			return
		}

		res, err := esClient.Ping()
		if err != nil {
			logrus.WithError(err).Errorln("Failed ping database")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), r.Context().Value(config.RequestID)).Stringify())
			return
		}

		if res.IsError() {
			logrus.WithError(err).Errorln("Failed ping database")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, cookbook.ErrorResponse(res.String(), r.Context().Value(config.RequestID)).Stringify())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, cookbook.SuccessResponse().Stringify())
	}
}

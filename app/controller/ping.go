package controller

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"

	"github.com/dynastymasra/goframe/config"

	"github.com/elastic/go-elasticsearch/v7"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dynastymasra/cookbook"
	"github.com/dynastymasra/cookbook/message"
	"github.com/sirupsen/logrus"
)

// Remove unused params
func Ping(db *gorm.DB, client *mongo.Client, esClient *elasticsearch.Client, neo4jDriver neo4j.Driver, redisClient *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(cookbook.HContentType, cookbook.HJSONTypeUTF8)
		w.Header().Set(cookbook.XRequestID, fmt.Sprintf("%v", r.Context().Value(config.RequestID)))

		log := logrus.WithFields(logrus.Fields{
			config.RequestID:    r.Context().Value(config.RequestID),
			config.JServiceName: config.ServiceName,
			config.JVersion:     config.Version,
		})

		if err := neo4jDriver.VerifyConnectivity(); err != nil {
			log.WithError(err).Errorln("Failed connect to Neo4J database")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), message.ErrDatabaseUnavailableCode).Stringify())
			return
		}

		dbClient, err := db.DB()
		if err != nil {
			log.WithError(err).Errorln("Failed get Postgres database")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), message.ErrDatabaseUnavailableCode).Stringify())
			return
		}

		if err := dbClient.Ping(); err != nil {
			log.WithError(err).Errorln("Failed connect to Postgres database")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), message.ErrDatabaseUnavailableCode).Stringify())
			return
		}

		if err := client.Ping(r.Context(), nil); err != nil {
			log.WithError(err).Errorln("Failed connect to Postgres database")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), message.ErrDatabaseUnavailableCode).Stringify())
			return
		}

		res, err := esClient.Ping()
		if err != nil {
			log.WithError(err).Errorln("Failed connect to elasticsearch database")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), message.ErrDatabaseUnavailableCode).Stringify())
			return
		}

		if res.IsError() {
			log.WithError(err).Errorln("Elasticsearch database connection has an error")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(res.String(), message.ErrDatabaseUnavailableCode).Stringify())
			return
		}

		if err := redisClient.Ping(r.Context()).Err(); err != nil {
			log.WithError(err).Errorln("Failed connect to redis database")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, cookbook.ErrorResponse(res.String(), message.ErrDatabaseUnavailableCode).Stringify())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, cookbook.SuccessResponse().Stringify())
	}
}

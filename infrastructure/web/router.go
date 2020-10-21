package web

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"

	"github.com/dynastymasra/cookbook/message"

	"github.com/dynastymasra/goframe/app/controller"
	"github.com/elastic/go-elasticsearch/v7"
	"gorm.io/gorm"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dynastymasra/goframe/config"

	"github.com/dynastymasra/cookbook/negroni/middleware"

	"github.com/urfave/negroni/v2"

	"github.com/dynastymasra/cookbook"
	"github.com/gorilla/mux"
)

type RouterInstance struct {
	PostgresDB  *gorm.DB
	Neo4JDriver neo4j.Driver
	MongoClient *mongo.Client
	EsClient    *elasticsearch.Client
	RedisClient *redis.Client
}

func (r *RouterInstance) Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true).UseEncodedPath()

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(cookbook.HContentType, cookbook.HJSONTypeUTF8)
		w.Header().Set(cookbook.HAccept, cookbook.HJSONTypeUTF8)
		w.WriteHeader(http.StatusNotFound)

		res := message.ErrEndpointNotFoundCode.ErrorMessage()
		fmt.Fprint(w, cookbook.FailResponse([]cookbook.JSON{{
			"title":   res.Title,
			"code":    res.Code,
			"message": res.Error.Error(),
		}}).Stringify())
	})

	router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(cookbook.HContentType, cookbook.HJSONTypeUTF8)
		w.Header().Set(cookbook.HAccept, cookbook.HJSONTypeUTF8)
		w.WriteHeader(http.StatusMethodNotAllowed)

		res := message.ErrMethodNotAllowedCode.ErrorMessage()
		fmt.Fprint(w, cookbook.FailResponse([]cookbook.JSON{{
			"title":   res.Title,
			"code":    res.Code,
			"message": res.Error.Error(),
		}}).Stringify())
	})

	commonHandlers := negroni.New(
		middleware.RequestID(config.RequestID),
	)

	// Probes
	router.Handle("/ping", commonHandlers.With(
		negroni.WrapFunc(controller.Ping(r.PostgresDB, r.MongoClient, r.EsClient, r.Neo4JDriver, r.RedisClient)),
	)).Methods(http.MethodGet, http.MethodHead)

	_ = router.PathPrefix("/v1/").Subrouter().UseEncodedPath()
	commonHandlers.Use(middleware.LogrusLog(config.ServiceName, config.RequestID))

	return router
}

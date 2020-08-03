package web

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/goframe/app/controller"
	"github.com/elastic/go-elasticsearch/v7"

	"github.com/neo4j/neo4j-go-driver/neo4j"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dynastymasra/goframe/config"

	"github.com/dynastymasra/cookbook/negroni/middleware"

	"github.com/urfave/negroni"

	"github.com/dynastymasra/cookbook"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

const DefaultResponseNotFound = "the requested resource doesn't exists"

type RouterInstance struct {
	PostgresDB  *gorm.DB
	Neo4JDriver neo4j.Driver
	MongoClient *mongo.Client
	EsClient    *elasticsearch.Client
}

func (r *RouterInstance) Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true).UseEncodedPath()

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, cookbook.FailResponse(&cookbook.JSON{
			"endpoint": DefaultResponseNotFound,
		}, nil).Stringify())
	})

	router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, cookbook.FailResponse(&cookbook.JSON{
			"method": DefaultResponseNotFound,
		}, nil).Stringify())
	})

	commonHandlers := negroni.New(
		middleware.RequestID(config.RequestID),
	)

	// Probes
	router.Handle("/ping", commonHandlers.With(
		negroni.WrapFunc(controller.Ping(r.PostgresDB, r.MongoClient, r.EsClient)),
	)).Methods(http.MethodGet, http.MethodHead)

	_ = router.PathPrefix("/v1/").Subrouter().UseEncodedPath()
	commonHandlers.Use(middleware.LogrusLog(config.ServiceName, config.RequestID))

	return router
}

package controller_test

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"github.com/neo4j/neo4j-go-driver/neo4j"

	"github.com/dynastymasra/goframe/app/controller"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dynastymasra/goframe/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type PingSuite struct {
	suite.Suite
	mg *mongo.Client
	db *gorm.DB
	es *elasticsearch.Client
	nj neo4j.Driver
	rd *redis.Client
}

func Test_PingSuite(t *testing.T) {
	suite.Run(t, new(PingSuite))
}

func (p *PingSuite) SetupSuite() {
	config.Load()
	config.SetupTestLogger()
}

func (p *PingSuite) SetupTest() {
	db, err := config.Postgres().Client()
	if err != nil {
		log.Fatalln(err)
	}
	mg, err := config.MongoDB().Client()
	if err != nil {
		log.Fatalln(err)
	}
	es, err := config.Elasticsearch().Client()
	if err != nil {
		log.Fatalln(err)
	}
	nj, err := config.Neo4J().Driver()
	if err != nil {
		log.Fatalln(err)
	}
	rd, err := config.Redis().Client()
	if err != nil {
		log.Fatalln(err)
	}

	p.db = db
	p.mg = mg
	p.es = es
	p.nj = nj
	p.rd = rd
}

func (p *PingSuite) Test_PingHandler_Success() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/ping", nil)
	ctx := context.WithValue(r.Context(), config.RequestID, uuid.NewV4().String())

	controller.Ping(p.db, p.mg, p.es, p.nj, p.rd)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusOK, w.Code)
}

func (p *PingSuite) Test_PingHandler_Failed() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/ping", nil)
	ctx := context.WithValue(r.Context(), config.RequestID, uuid.NewV4().String())

	p.nj.Close()
	controller.Ping(p.db, p.mg, p.es, p.nj, p.rd)(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusServiceUnavailable, w.Code)
}

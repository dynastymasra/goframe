package controller_test

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dynastymasra/goframe/app/controller"
	uuid "github.com/satori/go.uuid"

	"github.com/dynastymasra/goframe/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PingSuite struct {
	suite.Suite
}

func Test_PingSuite(t *testing.T) {
	suite.Run(t, new(PingSuite))
}

func (p *PingSuite) SetupSuite() {
	config.Load()
	config.SetupTestLogger()
}

func (p *PingSuite) SetupTest() {
	if _, err := config.Postgres().Client(); err != nil {
		log.Fatalln(err)
	}
	if _, err := config.MongoDB().Client(); err != nil {
		log.Fatalln(err)
	}
	if _, err := config.Elasticsearch().Client(); err != nil {
		log.Fatalln(err)
	}
	if _, err := config.Neo4J().Driver(); err != nil {
		log.Fatalln(err)
	}
	if _, err := config.Redis().Client(); err != nil {
		log.Fatalln(err)
	}
}

func (p *PingSuite) Test_PingHandler_Success() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/ping", nil)
	ctx := context.WithValue(r.Context(), config.RequestID, uuid.NewV4().String())

	controller.Ping()(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusOK, w.Code)
}

func (p *PingSuite) Test_PingHandler_Failed() {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/ping", nil)
	ctx := context.WithValue(r.Context(), config.RequestID, uuid.NewV4().String())

	p.nj.Close()
	controller.Ping()(w, r.WithContext(ctx))

	assert.Equal(p.T(), http.StatusServiceUnavailable, w.Code)
}

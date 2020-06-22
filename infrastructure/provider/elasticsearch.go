package provider

import (
	"errors"
	"net/http"
	"strings"

	"github.com/matryer/resync"

	"github.com/elastic/go-elasticsearch/v7"
)

var (
	esClient *elasticsearch.Client
	errEs    error
	runEs    resync.Once
)

type Elasticsearch struct {
	Address        string
	Username       string
	Password       string
	MaxConnPerHost int
	MaxIdlePerHost int
}

func (e Elasticsearch) Client() (*elasticsearch.Client, error) {
	addresses := strings.Split(e.Address, ",")

	config := elasticsearch.Config{
		Addresses: addresses,
		Username:  e.Username,
		Password:  e.Password,
		Transport: &http.Transport{
			MaxConnsPerHost:     e.MaxConnPerHost,
			MaxIdleConnsPerHost: e.MaxIdlePerHost,
		},
	}

	runEs.Do(func() {
		esClient, errEs = elasticsearch.NewClient(config)
	})

	res, err := esClient.Ping()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, errors.New(res.String())
	}

	return esClient, errEs
}

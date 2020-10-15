package config

import (
	"log"

	"github.com/dynastymasra/cookbook"
	"github.com/dynastymasra/cookbook/provider/elastic"
	"github.com/dynastymasra/cookbook/provider/mongo"
	"github.com/dynastymasra/cookbook/provider/neo4j"
	"github.com/dynastymasra/cookbook/provider/postgres"
	"github.com/spf13/viper"
)

type Config struct {
	serverPort string
	logger     LoggerConfig
	postgres   postgres.Config
	neo4j      neo4j.Config
	mongodb    mongo.Config
	elastic    elastic.Config
}

var config *Config

func Load() {
	viper.SetDefault(envServerPort, "8080")
	viper.SetDefault(envLoggerCaller, true)

	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../../")
	viper.AddConfigPath("../../../../")
	viper.SetConfigType("yaml")

	viper.ReadInConfig()

	config = &Config{
		serverPort: getString(envServerPort),
		logger: LoggerConfig{
			format: getString(envLoggerFormat),
			level:  getString(envLoggerLevel),
			caller: getBool(envLoggerCaller),
		},
		postgres: postgres.Config{
			Database:    getString(envPostgresDatabase),
			Host:        getString(envPostgresHost),
			Port:        getInt(envPostgresPort),
			Username:    getString(envPostgresUsername),
			Password:    getString(envPostgresPassword),
			Params:      getString(envPostgresParams),
			MaxIdleConn: getInt(envPostgresMaxIdleConn),
			MaxOpenConn: getInt(envPostgresMaxOpenConn),
			LogMode:     getInt(envPostgresLogLevel),
		},
		neo4j: neo4j.Config{
			Address:     getString(envNeo4JAddress),
			Username:    getString(envNeo4JUsername),
			Password:    getString(envNeo4JPassword),
			MaxConnPool: getInt(envNeo4JMaxConnPool),
			LogEnabled:  getBool(envNeo4JLogEnabled),
			LogLevel:    getInt(envNeo4JLogLevel),
		},
		mongodb: mongo.Config{
			Address:     getString(envMongoAddress),
			Username:    getString(envMongoUsername),
			Password:    getString(envMongoPassword),
			Database:    getString(envMongoDatabase),
			MaxPoolSize: getInt(envMongoMaxPool),
		},
		elastic: elastic.Config{
			Address:        getString(envElasticsearchAddress),
			Username:       getString(envElasticsearchUsername),
			Password:       getString(envElasticsearchPassword),
			MaxConnPerHost: getInt(envElasticsearchMaxConnPerHost),
			MaxIdlePerHost: getInt(envElasticsearchMaxIdlePerHost),
		},
	}
}

func ServerPort() string {
	return config.serverPort
}

func Logger() LoggerConfig {
	return config.logger
}

func Postgres() postgres.Config {
	return config.postgres
}

func Neo4J() neo4j.Config {
	return config.neo4j
}

func MongoDB() mongo.Config {
	return config.mongodb
}

func Elasticsearch() elastic.Config {
	return config.elastic
}

func getString(key string) string {
	value, err := cookbook.StringEnv(key)
	if err != nil {
		log.Fatalf("%v env key is not set", key)
	}
	return value
}

func getInt(key string) int {
	value, err := cookbook.IntEnv(key)
	if err != nil {
		log.Fatalf("%v env key is not set", key)
	}
	return value
}

func getBool(key string) bool {
	value, err := cookbook.BoolEnv(key)
	if err != nil {
		log.Fatalf("%v env key is not set", key)
	}
	return value
}

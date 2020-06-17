package config

import (
	"log"

	"github.com/dynastymasra/goframe/infrastructure/provider"

	"github.com/dynastymasra/cookbook"
	"github.com/spf13/viper"
)

type Config struct {
	serverPort string
	logger     LoggerConfig
	postgres   provider.Postgres
	neo4j      provider.Neo4J
}

var config *Config

func Load() {
	viper.SetDefault(envServerPort, "8080")
	viper.SetDefault(envLoggerCaller, true)

	viper.AutomaticEnv()

	viper.SetConfigName("application")
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
		postgres: provider.Postgres{
			DatabaseName: getString(envPostgresDBName),
			Address:      getString(envPostgresAddress),
			Username:     getString(envPostgresUsername),
			Password:     getString(envPostgresPassword),
			MaxIdleConn:  getInt(envPostgresMaxIdleConn),
			MaxOpenConn:  getInt(envPostgresMaxOpenConn),
			LogEnabled:   getBool(envPostgresLogEnabled),
		},
		neo4j: provider.Neo4J{
			Address:     getString(envNeo4JAddress),
			Username:    getString(envNeo4JUsername),
			Password:    getString(envNeo4JPassword),
			MaxConnPool: getInt(envNeo4JMaxConnPool),
			Encrypted:   getBool(envNeo4JEncrypted),
			LogEnabled:  getBool(envNeo4JLogEnabled),
			LogLevel:    getInt(envNeo4JLogLevel),
		},
	}
}

func ServerPort() string {
	return config.serverPort
}

func Logger() LoggerConfig {
	return config.logger
}

func Postgres() provider.Postgres {
	return config.postgres
}

func Neo4J() provider.Neo4J {
	return config.neo4j
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

package config

import (
	"log"

	"github.com/dynastymasra/cookbook"
	"github.com/spf13/viper"
)

type Config struct {
	serverPort string
	logger     LoggerConfig
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
	}
}

func ServerPort() string {
	return config.serverPort
}

func Logger() LoggerConfig {
	return config.logger
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

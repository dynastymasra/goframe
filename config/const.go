package config

const (
	ServiceName = "Goframe"
	Version     = "0.1.0"

	// Service port address
	envServerPort = "SERVER_PORT"
	envGRPCPort   = "GRPC_PORT"

	// Logger config
	envLoggerFormat = "LOGGER_FORMAT"
	envLoggerLevel  = "LOGGER_LEVEL"
	envLoggerCaller = "LOGGER_CALLER_ENABLED"

	// Postgres config
	envPostgresDBName      = "POSTGRES_DATABASE_NAME"
	envPostgresAddress     = "POSTGRES_ADDRESS"
	envPostgresUsername    = "POSTGRES_USERNAME"
	envPostgresPassword    = "POSTGRES_PASSWORD"
	envPostgresMaxIdleConn = "POSTGRES_MAX_IDLE_CONN"
	envPostgresMaxOpenConn = "POSTGRES_MAX_OPEN_CONN"
	envPostgresLogEnabled  = "POSTGRES_LOG_ENABLED"
)

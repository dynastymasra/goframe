package config

const (
	ServiceName = "Goframe"
	Version     = "0.1.0"
	RequestID   = "requestId"

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

	// Neo4J config
	envNeo4JAddress     = "NEO4J_ADDRESS"
	envNeo4JUsername    = "NEO4J_USERNAME"
	envNeo4JPassword    = "NEO4J_PASSWORD"
	envNeo4JMaxConnPool = "NEO4J_MAX_CONN_POOL"
	envNeo4JEncrypted   = "NEO4J_ENCRYPTED"
	envNeo4JLogEnabled  = "NEO4J_LOG_ENABLED"
	envNeo4JLogLevel    = "NEO4J_LOG_LEVEL"

	// MongoDB config
	envMongoFormat   = "MONGO_FORMAT"
	envMongoAddress  = "MONGO_ADDRESS"
	envMongoDatabase = "MONGO_DATABASE"
	envMongoUsername = "MONGO_USERNAME"
	envMongoPassword = "MONGO_PASSWORD"
)

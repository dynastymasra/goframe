package config

const (
	ServiceName  = "Goframe"
	Version      = "0.1.0"
	RequestID    = "requestId"
	JServiceName = "service"
	JVersion     = "version"

	// Service port address
	envServerPort = "SERVER_PORT"

	// Logger config
	envLoggerFormat = "LOGGER_FORMAT"
	envLoggerLevel  = "LOGGER_LEVEL"
	envLoggerCaller = "LOGGER_CALLER_ENABLED"

	// Postgres config
	envPostgresDatabase     = "POSTGRES_DATABASE"
	envPostgresHost         = "POSTGRES_HOST"
	envPostgresPort         = "POSTGRES_PORT"
	envPostgresUsername     = "POSTGRES_USERNAME"
	envPostgresPassword     = "POSTGRES_PASSWORD"
	envPostgresParams       = "POSTGRES_PARAMS"
	envPostgresMaxIdleConn  = "POSTGRES_MAX_IDLE_CONN"
	envPostgresMaxOpenConn  = "POSTGRES_MAX_OPEN_CONN"
	envPostgresLogLevel     = "POSTGRES_LOG_LEVEL"
	envPostgresDebugEnabled = "POSTGRES_DEBUG_ENABLED"

	// Neo4J config
	envNeo4JAddress         = "NEO4J_ADDRESS"
	envNeo4JUsername        = "NEO4J_USERNAME"
	envNeo4JPassword        = "NEO4J_PASSWORD"
	envNeo4JMaxConnPool     = "NEO4J_MAX_CONN_POOL"
	envNeo4JMaxConnLifetime = "NEO4J_MAX_CONN_LIFETIME"
	envNeo4JLogEnabled      = "NEO4J_LOG_ENABLED"
	envNeo4JLogLevel        = "NEO4J_LOG_LEVEL"
	envNeo4JVerifyHostname  = "NEO4J_VERIFY_HOSTNAME"

	// MongoDB config
	envMongoAddress  = "MONGO_ADDRESS"
	envMongoDatabase = "MONGO_DATABASE"
	envMongoUsername = "MONGO_USERNAME"
	envMongoPassword = "MONGO_PASSWORD"
	envMongoMaxPool  = "MONGO_MAX_POOL"

	// Elasticsearch config
	envElasticsearchAddress        = "ELASTICSEARCH_ADDRESS"
	envElasticsearchUsername       = "ELASTICSEARCH_USERNAME"
	envElasticsearchPassword       = "ELASTICSEARCH_PASSWORD"
	envElasticsearchMaxConnPerHost = "ELASTICSEARCH_MAX_CONN_PER_HOST"
	envElasticsearchMaxIdlePerHost = "ELASTICSEARCH_MAX_IDLE_PER_HOST"

	// Redis config
	envRedisAddress     = "REDIS_ADDRESS"
	envRedisPassword    = "REDIS_PASSWORD"
	envRedisDatabase    = "REDIS_DATABASE"
	envRedisPoolSize    = "REDIS_POOL_SIZE"
	envRedisMinIdleConn = "REDIS_MIN_IDLE_CONN"
)

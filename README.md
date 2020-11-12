# Goframe

[![Go](https://img.shields.io/badge/go-1.14-00E5E6.svg)](https://golang.org/)
[![Postgres](https://img.shields.io/badge/postgres-12.3-326590.svg)](https://www.postgresql.org/)
[![MongoDB](https://img.shields.io/badge/mongodb-4.2.8-139B50.svg)](https://www.mongodb.com/)
[![Neo4J](https://img.shields.io/badge/neo4j-4.1.3-3A8B9F.svg)](https://neo4j.com/)
[![Elasticsearch](https://img.shields.io/badge/elasticsearch-7.7.0-F4BE1A.svg)](https://www.elastic.co/elasticsearch/)
[![Redis](https://img.shields.io/badge/redis-6.0.8-C93D2E.svg)](https://redis.io/)
[![Seabolt](https://img.shields.io/badge/seabolt-1.7.4-2885E4.svg)](https://github.com/neo4j-drivers/seabolt)
[![Docker](https://img.shields.io/badge/docker-19.03-2885E4.svg)](https://www.docker.com/)

Go skeleton (MongoDB, Neo4J, Postgres), Remove unused or unnecessary function or code

## Libraries

Use [Go Module](https://blog.golang.org/using-go-modules) for install all dependencies required this application.

If use private repository dependency please make use GIT Private key created `ssh-keygen -t rsa -P "" -C "email" -m PEM`

## How To Run and Deploy

Before run this service. Make sure all requirements dependencies has been installed likes **Golang, Docker, and database**

### Structure

This service tries to implement the clean architecture.

```
├── app (Interface Adapters) 
│   ├── controller
│   └── presenter
│── config (Application config)
│── console (Terminal console)
│── domain (Enterprise Business Rules)
│   ├── repository (Application Business Rules)
│   └── service (Application Business Rules)
│── infrastructure (Frameworks & Drivers)
│   ├── provider (Drivers)
│   └── web (Web Application)
│── migration (Migration Files)
│── test (Test Helper & Mocks)
└── main.go (Main Application)
```

### Local

Use command go ```go run main.go``` in root folder for run this application.

### Migration

- ```go run main.go migrate:run``` Command used to run the migration files.
- ```go run main.go migrate:rollback``` Command used to rollback the migration.
- ```go run main.go migrate:create``` Command used to create a new migration file.

### Docker

**goframe** uses docker multi stages build, minimal docker version is **17.05**. If docker already installed use command.

Insert arguments `GIT_PRIVATE_KEY` and `GOPRIVATE` for create an image

This command will build the images.
```bash
docker build -f Dockerfile  --build-arg GIT_PRIVATE_KEY=$GIT_PRIVATE_KEY --build-arg GOPRIVATE=$GOPRIVATE -t goframe:$(VERSION) .
```

To run service use this command
```bash
docker run --name goframe -d -e ADDRESS=:8080 -e <environment> $(IMAGE):$(VERSION)
```

## Test

### Local

For run unit test, from root project you can go to folder or package and execute command
```bash
go test -v -cover -coverprofile=coverage.out -covermode=set
go tool cover -html=coverage.out
```
`go tool` will generate GUI for test coverage. Available package or folder can be tested.

- `/app/controller`
- `/app/presenter`
- `/config`
- `/domain`

### Docker

Run docker-compose to run test
```
docker-compose -p "$(PROJECT_NAME)" -f docker/docker-compose.yaml build app
docker-compose -p "$(PROJECT_NAME)" -f docker/docker-compose.yaml run --rm app test
docker-compose -p "$(PROJECT_NAME)" -f docker/docker-compose.yaml down
```

## Environment Variables

Environment variables for Development use **config-example.yaml**, Change the file name to **config.yaml**

+ `SERVER_PORT` - Port address used by service, default is `8080`
+ `LOGGER_LEVEL` - Log level(debug, info, error, warn, etc)
+ `LOGGER_CALLER_ENABLED` - Enabled log report caller.
+ `LOGGER_FORMAT` - Format specific for log.
  - `text` - Log format will become standard text output, this used for development
  - `json` - Log format will become *JSON* format, usually used for production
+ `NEO4J_ADDRESS` - Neo4J database address `bolt+routing://<host>:<port>`
  - `bolt+routing://` - Used with causal cluster
  - `bolt://` - Used with single server
+ `NEO4J_USERNAME` - Neo4J database username
+ `NEO4J_PASSWORD` - Neo4J database password
+ `NEO4J_MAX_CONN_POOL` - Neo4j maximum number of connections per URL to allow on this driver
+ `NEO4J_MAX_CONN_LIFETIME` - Maximum connection life time on pooled connections. Values less than or equal to 0 disables the lifetime check (in Minutes)
+ `NEO4J_LOG_ENABLED` - Neo4J database log enabled (`true`/`false`)
+ `NEO4J_LOG_LEVEL` Neo4J type that default logging implementations
  - `1` - Level error
  - `2` - Level warning
  - `3` - Level info
  - `4` - Level debug
+ `POSTGRES_ADDRESS` - Database hostname
+ `POSTGRES_DATABASE` - Database name
+ `POSTGRES_USERNAME` - Database username
+ `POSTGRES_PASSWORD` - Database Password
+ `POSTGRES_PARAMS` - Database params, use space if more than one param (sslmode=disable)
+ `POSTGRES_MAX_OPEN_CONN` - Database max open connection
+ `POSTGRES_MAX_IDLE_CONN` - Database max idle connection
+ `POSTGRES_LOG_LEVEL` - Database log level, default is 2
    - `1` - Silent
    - `2` - Error
    - `3` - Warn
    - `4` - Info
+ `MONGO_ADDRESS` - MongoDB address with `mongodb://<host>:<port>`, `mongodb+srv://<host>:<port>`, and `"mongodb://<host>:<port>/?replicaSet=<replica>&connect=direct"` read the documentation for more information
+ `MONGO_DATABASE` - MongoDB database name
+ `MONGO_USERNAME` - MongoDB username
+ `MONGO_PASSWORD` - MongoDB password
+ `MONGO_MAX_POOL` - MongoDB max pool connection
+ `ELASTICSEARCH_ADDRESS` - Elasticsearch urls, use `,` to separate urls `http://localhost:9200,http://localhost:9201`
+ `ELASTICSEARCH_USERNAME` - Elasticsearch username
+ `ELASTICSEARCH_PASSWORD` - Elasticsearch password
+ `ELASTICSEARCH_MAX_CONN_PER_HOST` - Elasticsearch max connection per hosts
+ `ELASTICSEARCH_MAX_IDLE_PER_HOST` - Elasticsearch max idle connection per hosts
+ `REDIS_ADDRESS` - Redis address `localhost:6379`
+ `REDIS_PASSWORD` - Redis password
+ `REDIS_DATABASE` - Redis database
+ `REDIS_POOL_SIZE` - Redis pool size
+ `REDIS_MIN_IDLE_CONN` - Redis min idle connection

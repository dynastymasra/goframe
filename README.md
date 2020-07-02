# Goframe

[![Go](https://img.shields.io/badge/go-1.14-00E5E6.svg)](https://golang.org/)
[![Postgres](https://img.shields.io/badge/postgres-12.3-326590.svg)](https://www.postgresql.org/)
[![MongoDB](https://img.shields.io/badge/mongodb-4.2.8-139B50.svg)](https://www.mongodb.com/)
[![Elasticsearch](https://img.shields.io/badge/elasticsearch-4.2.8-F4BE1A.svg)](https://www.elastic.co/elasticsearch/)
[![Docker](https://img.shields.io/badge/docker-19.03-2885E4.svg)](https://www.docker.com/)

Go skeleton (MongoDB, Neo4J, Postgres), Remove unused or unnecessary function or code

## Libraries

Use [Go Module](https://blog.golang.org/using-go-modules) for install all dependencies required this application.

If use private repository dependency please make use GIT Private key created `ssh-keygen -t rsa -P "" -C "email" -m PEM`

## How To Run and Deploy

Before run this service. Make sure all requirements dependencies has been installed likes **Golang, Docker, and database**

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

For run unit test, from root project you can go to folder or package and execute command
```bash
go test -v -cover -coverprofile=coverage.out -covermode=set
go tool cover -html=coverage.out
```
`go tool` will generate GUI for test coverage. Available package or folder can be tested

- `/infrastructure/web/handler`

## Environment Variables

+ `SERVER_PORT` - Address application is used default is `8080`
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
+ `NEO4J_ENCRYPTED` - Neo4J whether to turn on/off TLS encryption (`true`/`false`)
+ `NEO4J_LOG_ENABLED` - Neo4J database log enabled (`true`/`false`)
+ `NEO4J_LOG_LEVEL` Neo4J type that default logging implementations use for available default `0`
  - `0` - Doesn't generate any output
  - `1` - Level error
  - `2` - Level warning
  - `3` - Level info
  - `4` - Level debug
+ `POSTGRES_ADDRESS` - Database hostname
+ `POSTGRES_DATABASE` - Database name
+ `POSTGRES_USERNAME` - Database username
+ `POSTGRES_PASSWORD` - Database Password
+ `POSTGRES_LOG_ENABLED` - Database log enabled, value `true` or `false`
+ `POSTGRES_MAX_OPEN_CONN` - Database max open connection
+ `POSTGRES_MAX_IDLE_CONN` - Database max idle connection
+ `MONGO_ADDRESS` - MongoDB address with `mongodb://<host>:<port>` or `mongodb+srv://<host>:<port>`
+ `MONGO_DATABASE` - MongoDB database name
+ `MONGO_USERNAME` - MongoDB username
+ `MONGO_PASSWORD` - MongoDB password
+ `MONGO_MAX_POOL` - MongoDB max pool connection
+ `ELASTICSEARCH_ADDRESS` - Elasticsearch urls, use `,` to separate urls `http://localhost:9200,http://localhost:9201`
+ `ELASTICSEARCH_USERNAME` - Elasticsearch username
+ `ELASTICSEARCH_PASSWORD` - Elasticsearch password
+ `ELASTICSEARCH_MAX_CONN_PER_HOST` - Elasticsearch max connection per hosts
+ `ELASTICSEARCH_MAX_IDLE_PER_HOST` - Elasticsearch max idle connection per hosts

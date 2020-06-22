FROM alpine:3.11 AS seabolt

# install dependecies
RUN set -ex \
    && apk add --update git openssl-dev cmake make g++ openssl-libs-static \
    && rm -rf /var/cache/apk/*

# create seabolt package Neo4J dependency
RUN git clone -b v1.7.4 https://github.com/neo4j-drivers/seabolt.git /seabolt
WORKDIR /seabolt/build
RUN cmake -D CMAKE_BUILD_TYPE=Release -D CMAKE_INSTALL_LIBDIR=lib .. && cmake --build . --target package

############################## SEPARATOR ##############################

FROM golang:1.14-alpine AS builder

ARG GIT_PRIVATE_KEY
ARG GOPRIVATE

# install dependecies
RUN set -ex \
    && apk add --update git curl openssh-client ca-certificates pkgconfig g++ \
    && rm -rf /var/cache/apk/*

# install seabolt Neo4J dependency
COPY --from=seabolt /seabolt/build/dist-package/seabolt*.tar.gz /tmp
RUN tar zxf /tmp/seabolt*.tar.gz --strip-components=1 -C /

WORKDIR /go/src/github.com/dynastymasra/goframe
COPY go.mod go.sum git-creds.sh ./
RUN ./git-creds.sh gitlab.com
RUN go mod download

## build linux app source code
COPY . ./
RUN GOOS=linux go build -tags=main -o goframe

############################## SEPARATOR ##############################

FROM alpine:3.11

RUN set -ex \
    && apk add --update bash ca-certificates tzdata \
    && rm -rf /var/cache/apk/*

# install seabolt Neo4J dependency
COPY --from=seabolt /seabolt/build/dist-package/seabolt*.tar.gz /tmp
RUN tar zxf /tmp/seabolt*.tar.gz --strip-components=1 -C /

# app
WORKDIR /app
RUN mkdir migration
COPY --from=builder /go/src/github.com/dynastymasra/goframe/goframe /app/

# remove seabolt package
RUN rm -rf /tmp

## runtime configs
EXPOSE 8080

ENTRYPOINT ["./goframe"]
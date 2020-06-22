package provider

import (
	"fmt"

	"github.com/matryer/resync"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var (
	neo4jDriver neo4j.Driver
	errNeo4J    error
	runNeo4J    resync.Once
)

type Neo4J struct {
	Address     string
	Username    string
	Password    string
	MaxConnPool int
	Encrypted   bool
	LogEnabled  bool
	LogLevel    int
}

func (n Neo4J) Driver() (neo4j.Driver, error) {
	url := fmt.Sprintf("%s", n.Address)
	auth := neo4j.BasicAuth(n.Username, n.Password, "")

	runNeo4J.Do(func() {
		neo4jDriver, errNeo4J = neo4j.NewDriver(url, auth, func(config *neo4j.Config) {
			config.Encrypted = n.Encrypted
			config.MaxConnectionPoolSize = n.MaxConnPool
			if n.LogEnabled {
				config.Log = neo4j.ConsoleLogger(neo4j.LogLevel(n.LogLevel))
			}
		})
	})

	return neo4jDriver, errNeo4J
}

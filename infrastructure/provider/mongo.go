package provider

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/matryer/resync"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client   *mongo.Client
	errMongo error
	runMongo resync.Once
)

type MongoDB struct {
	Format   string
	Address  string
	Username string
	Password string
	Database string
}

func (m *MongoDB) Client() (*mongo.Client, error) {
	url := fmt.Sprintf("%s://%s:%s@%s/%s", m.Format, m.Username, m.Password, m.Address, m.Database)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	runMongo.Do(func() {
		client, errMongo = mongo.NewClient(options.Client().ApplyURI(url))

		errMongo = client.Connect(ctx)

		errMongo = client.Ping(ctx, nil)
	})

	return client, errMongo
}

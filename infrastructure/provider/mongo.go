package provider

import (
	"context"

	"github.com/matryer/resync"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
	errMongo    error
	runMongo    resync.Once
)

type MongoDB struct {
	Address     string
	Username    string
	Password    string
	Database    string
	MaxPoolSize int
}

func (m MongoDB) Client() (*mongo.Client, error) {
	auth := options.Credential{
		AuthSource:  m.Database,
		Username:    m.Password,
		Password:    m.Password,
		PasswordSet: true,
	}

	runMongo.Do(func() {
		mongoClient, errMongo = mongo.NewClient(options.Client().
			SetAuth(auth).
			SetMaxPoolSize(uint64(m.MaxPoolSize)).
			ApplyURI(m.Address))

		errMongo = mongoClient.Connect(context.Background())
	})

	if err := mongoClient.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return mongoClient, errMongo
}

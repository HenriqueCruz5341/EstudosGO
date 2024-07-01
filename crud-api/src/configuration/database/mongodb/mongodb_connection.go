package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MONGODB_URL     = "MONGODB_URL"
	MONGODB_USER    = "MONGODB_USER"
	MONGODB_PASS    = "MONGODB_PASS"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongoDbUrl := os.Getenv(MONGODB_URL)
	mongoDbUser := os.Getenv(MONGODB_USER)
	mongoDbPass := os.Getenv(MONGODB_PASS)
	mongoDbUserDb := os.Getenv(MONGODB_USER_DB)
	credential := options.Credential{
		Username: mongoDbUser,
		Password: mongoDbPass,
	}
	opts := options.Client().ApplyURI(mongoDbUrl).SetAuth(credential)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client.Database(mongoDbUserDb), nil
}

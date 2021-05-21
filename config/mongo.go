package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*10)
}

func MongoConnection(dbName, uri string, minPool, maxPool, maxIdle uint64) (*mongo.Database, error) {
	opts := options.Client().
		ApplyURI(uri).
		SetMinPoolSize(minPool).
		SetMaxPoolSize(maxPool).
		SetMaxConnIdleTime(time.Second * time.Duration(maxIdle))

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	ctx, cancel := MongoContext()
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database(dbName), nil
}

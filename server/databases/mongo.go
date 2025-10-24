package databases

import (
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Client, error) {
	username := "root"
	password := "password"
	host := "mongodb"
	port := 27017

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", username, password, host, port)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}

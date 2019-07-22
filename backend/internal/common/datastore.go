package common

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

// Connect to MongoDB
func Connect(ctx context.Context, host string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%v", host)))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Connected to MongoDB (%v)", host)
	return client, nil
}

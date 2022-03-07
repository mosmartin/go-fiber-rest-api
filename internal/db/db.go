package db

import (
	"context"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client

var mongoOnce sync.Once

var clientInstanceErr error

type Collection string

const (
	ProductsCollection Collection = "products"
)

func GetMongoClient() (*mongo.Client, error) {
	// Initialize the client once (singleton)
	mongoOnce.Do(func() {
		// clientInstance = getMongoClient()
		clientOPtions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))

		client, err := mongo.Connect(context.TODO(), clientOPtions)
		if err != nil {
			clientInstanceErr = err
			return
		}

		clientInstance = client

		clientInstanceErr = err
	})

	return clientInstance, clientInstanceErr
}

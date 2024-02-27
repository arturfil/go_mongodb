package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// Connect to MongoDB
func ConnectToMongo() (*mongo.Client, error) {
	// MongoDB connection string
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	username := os.Getenv("MONGO_DB_USERNAME")
	password := os.Getenv("MONGO_DB_PASSWORD")

	// set username and password
	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to mongo!")

	return client, nil
}

func GetCollectionPointer() *mongo.Collection {
	return collection
}

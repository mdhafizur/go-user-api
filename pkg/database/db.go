// Package db pkg/database/db.go
package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database

// Initialize initializes the database connection.
func Initialize(dbURL ...string) {
	var actualDBURL string

	if len(dbURL) > 0 && dbURL[0] != "" {
		// If dbURL is provided, use it
		actualDBURL = dbURL[0]
	} else {
		// If dbURL is not provided, fetch from the environment variable
		actualDBURL = os.Getenv("DB_URL")
		if actualDBURL == "" {
			log.Fatal("DB_URL environment variable not set")
		}
	}

	// Get the MongoDB database name from the environment variable
	mongoDBName := os.Getenv("MONGO_INITDB_DATABASE")
	// Set a default value if the environment variable is not set
	if mongoDBName == "" {
		mongoDBName = "testdb"
		log.Fatal("MONGO_INITDB_DATABASE environment variable not set")
	}

	clientOptions := options.Client().ApplyURI(actualDBURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB: ", err)
	}

	fmt.Println("Connected to MongoDB")

	Client = client

	Database = client.Database(mongoDBName)
}

// Disconnect closes the database connection.
func Disconnect() {
	if Client != nil {
		err := Client.Disconnect(context.Background())
		if err != nil {
			log.Fatal("Error disconnecting from MongoDB: ", err)
		}
		fmt.Println("Disconnected from MongoDB")
	}
}

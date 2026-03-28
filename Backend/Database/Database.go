package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var Client *mongo.Client

func init() {
	// The init function runs automatically when the package is imported
	Client = DBInstance()
}

func DBInstance() *mongo.Client {
	// Load .env here to be 100% sure it's available
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("MONGO_URL")
	if uri == "" {
		log.Fatal("MONGO_URL is not defined in .env")
	}

	// FIX: Use 'uri' directly, NOT os.Getenv(uri)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("MongoDB Connection Error: ", err)
	}

	// Verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not ping MongoDB: ", err)
	}

	return client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "assignment_db" // Fallback
	}
	return client.Database(dbName).Collection(collectionName)
}
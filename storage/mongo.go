package storage

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Initialize MongoDB with Atlas connection
func InitMongo() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Atlas connection string
	clientOptions := options.Client().ApplyURI(os.Getenv("SERVICE_MONGODB_URL"))

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("❌ Mongo Connect Error:", err)
		return err
	}

	// Ping database
	if err := client.Ping(ctx, nil); err != nil {
		log.Println("❌ Mongo Ping Error:", err)
		return err
	}

	log.Println("✅ MongoDB Atlas Connected Successfully")
	return nil
}

// Return Mongo client
func GetMongo() *mongo.Client {
	return client
}

// Get database
func GetDatabase(dbName string) *mongo.Database {
	return client.Database(dbName)
}

// Get collection inside fitpro database
func GetCollection(collectionName string) *mongo.Collection {
	return client.Database("fitpro").Collection(collectionName)
}

// Disconnect cleanly
func CloseMongo() {
	if client != nil {
		_ = client.Disconnect(context.Background())
	}
}
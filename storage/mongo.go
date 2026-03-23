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
	log.Println("⏳ Connecting to MongoDB...")
	connectCtx, cancelConnect := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelConnect()

	uri := os.Getenv("SERVICE_MONGODB_URL")
	if uri == "" {
		log.Fatal("SERVICE_MONGODB_URL is not set in .env")
	}

	clientOptions := options.Client().ApplyURI(uri)

	var err error
	client, err = mongo.Connect(connectCtx, clientOptions)
	if err != nil {
		log.Println("❌ Mongo Connect Error:", err)
		return err
	}

	log.Println("⏳ Pinging MongoDB...")
	pingCtx, cancelPing := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelPing()

	if err := client.Ping(pingCtx, nil); err != nil {
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
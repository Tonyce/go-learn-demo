package common

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database ...
type Database struct {
	*mongo.Client
}

var mongoClient *mongo.Client

// InitMongoDB ...
func init() {
	fmt.Println("mongoClient init")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/test"))
	if err != nil {
		log.Fatal(err)
	}
	mongoClient = client
	// defer client.Disconnect(ctx)
}

// GetMongoClient ...
func GetMongoClient() *mongo.Client {
	return mongoClient
}

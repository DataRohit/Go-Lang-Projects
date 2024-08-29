package dbConfig

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017"

var client *mongo.Client
var database *mongo.Database

func InitDB() error {
	log.Println("Connecting to database...")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return err
	}

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	database = client.Database("users")

	log.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return nil
}

func GetCollection(name string) *mongo.Collection {
	return database.Collection(name)
}

func Disconnect(ctx context.Context) error {
	return client.Disconnect(ctx)
}

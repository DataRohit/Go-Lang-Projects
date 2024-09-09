package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseInstance() *mongo.Client {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error loading the .env file")
	}

	connectionURI := os.Getenv("MONGODB_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb!")

	return client
}

var Client *mongo.Client = DatabaseInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	databaseName := os.Getenv("DATABASE_NAME")

	var collection *mongo.Collection = client.Database(databaseName).Collection(collectionName)
	return collection
}

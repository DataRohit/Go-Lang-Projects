package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionURI string = os.Getenv("MONGODB_URI")
var databaseName string = os.Getenv("DATABASE_NAME")

func DatabaseInstance() *mongo.Client {
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
	var collection *mongo.Collection = client.Database(databaseName).Collection(collectionName)
	return collection
}

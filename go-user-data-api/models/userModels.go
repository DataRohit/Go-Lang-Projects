package userModels

import (
	"context"
	"fmt"
	"log"

	dbConfig "github.com/datarohit/go-user-data-api/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Gender string             `json:"gender" bson:"gender"`
	Age    int                `json:"age" bson:"age"`
}

func GetAllUsers(ctx context.Context) ([]User, error) {
	var users []User

	collection := dbConfig.GetCollection("users")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("Failed to execute find query: %w", err)
	}
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Printf("Warning: Failed to close cursor: %v", err)
		}
	}()

	for cursor.Next(ctx) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, fmt.Errorf("Failed to decode user document: %w", err)
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("Cursor encountered an error: %w", err)
	}

	return users, nil
}

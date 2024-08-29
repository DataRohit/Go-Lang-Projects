package userModels

import (
	"context"
	"fmt"
	"log"

	dbConfig "github.com/datarohit/go-user-data-api/config"
	"github.com/datarohit/go-user-data-api/schemas"
	"github.com/datarohit/go-user-data-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllUsers(ctx context.Context) ([]schemas.User, error) {
	var users []schemas.User

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
		var user schemas.User
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

func GetUserByID(ctx context.Context, id string) (schemas.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid user ID format: %v", err)
		return schemas.User{}, fmt.Errorf("Invalid user ID format: %w", err)
	}

	collection := dbConfig.GetCollection("users")

	var user schemas.User

	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		log.Printf("Error finding user with ID %s: %v", id, err)
		return schemas.User{}, fmt.Errorf("Failed to find user with ID %s: %w", id, err)
	}

	return user, nil
}

func CreateUser(ctx context.Context, user schemas.User) (schemas.User, error) {
	if err := utils.ValidateUser(user); err != nil {
		log.Printf("User validation failed: %v", err)
		return schemas.User{}, fmt.Errorf("Validation error: %w", err)
	}

	user.Id = primitive.NewObjectID()

	collection := dbConfig.GetCollection("users")

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Error creating user in the database: %v", err)
		return schemas.User{}, fmt.Errorf("Failed to create user: %w", err)
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.Id = oid
	} else {
		log.Printf("Unexpected type for inserted ID: %v", result.InsertedID)
		return schemas.User{}, fmt.Errorf("Failed to retrieve the inserted ID")
	}

	log.Printf("User created successfully with ID: %s", user.Id.Hex())
	return user, nil
}

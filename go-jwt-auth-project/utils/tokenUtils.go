package utils

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/datarohit/go-jwt-auth-project/database"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	UID       string
	UserType  string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.DatabaseInstance(), "users")
var JWT_SECRET string = os.Getenv("JWT_SECRET")

func GenerateAllTokens(email string, firstName string, lastName string, userType, uid string) (signedAccessToken string, signedRefreshToken string, err error) {
	accessClaims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		UID:       uid,
		UserType:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(JWT_SECRET))
	if err != nil {
		log.Panic(err)
		return
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(JWT_SECRET))
	if err != nil {
		log.Panic(err)
		return
	}

	return accessToken, refreshToken, err
}

func ValidateToken(signedAccessToken string) (accessClaims *SignedDetails, msg string) {
	accessToken, err := jwt.ParseWithClaims(
		signedAccessToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_SECRET), nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}

	accessClaims, ok := accessToken.Claims.(*SignedDetails)
	if !ok {
		msg = "the provided access token is invalid"
		return
	}

	if accessClaims.ExpiresAt < time.Now().Local().Unix() {
		msg = "the provided access token is expired"
		return
	}

	return accessClaims, msg
}

func UpdateAllTokens(signedAccessToken string, signedRefreshToken string, userId string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var updatedObj primitive.D

	updatedObj = append(updatedObj, bson.E{Key: "accessToken", Value: signedAccessToken})
	updatedObj = append(updatedObj, bson.E{Key: "refreshToken", Value: signedRefreshToken})

	updatedAt, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		log.Panic(err)
		return
	}

	updatedObj = append(updatedObj, bson.E{Key: "updatedAt", Value: updatedAt})

	upsert := true
	filter := bson.M{"userId": userId}
	opts := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err = userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{{Key: "$set", Value: updatedObj}},
		&opts,
	)
	defer cancel()

	if err != nil {
		log.Panic(err)
		return
	}

	return
}

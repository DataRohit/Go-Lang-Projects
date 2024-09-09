package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/datarohit/go-jwt-auth-project/database"
	"github.com/datarohit/go-jwt-auth-project/models"
	"github.com/datarohit/go-jwt-auth-project/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
var validate = validator.New()

func SingUp() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := ginCtx.BindJSON(&user); err != nil {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking for the email"})
			return
		}

		if count > 0 {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "this email already exists"})
			return
		}

		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		if err != nil {
			log.Panic(err)
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking for the phone number"})
			return
		}

		if count > 0 {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "this phone number already exists"})
			return
		}

		password := utils.HashPassword(*user.Password)
		user.Password = &password
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		user.Id = primitive.NewObjectID()
		user.UserId = user.Id.Hex()

		accessToken, refreshToken, _ := utils.GenerateAllTokens(*user.Email, *user.FirstName, *user.LastName, *user.UserType, user.UserId)
		user.AccessToken = &accessToken
		user.RefreshToken = &refreshToken

		result, err := userCollection.InsertOne(ctx, user)
		if err != nil {
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "user item was not created"})
			return
		}

		ginCtx.JSON(http.StatusOK, result)
	}
}

func Login() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		var foundUser models.User

		if err := ginCtx.BindJSON(&user); err != nil {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		if err != nil {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect"})
			return
		}

		passwordIsValid, msg := utils.VerifyPassword(*user.Password, *foundUser.Password)
		if !passwordIsValid {
			ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			return
		}

		if foundUser.Email == nil {
			ginCtx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		accessToken, refreshToken, _ := utils.GenerateAllTokens(*foundUser.Email, *foundUser.FirstName, *foundUser.LastName, *foundUser.UserType, foundUser.UserId)

		utils.UpdateAllTokens(accessToken, refreshToken, foundUser.UserId)

		ginCtx.JSON(http.StatusOK, gin.H{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		})
	}
}

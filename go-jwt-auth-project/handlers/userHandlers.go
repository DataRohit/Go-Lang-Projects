package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/datarohit/go-jwt-auth-project/models"
	"github.com/datarohit/go-jwt-auth-project/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllUsers() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		if err := utils.CheckUserType(ginCtx, "ADMIN"); err != nil {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		recordsPerPage, err := strconv.Atoi(ginCtx.Query("recordsPerPage"))
		if err != nil || recordsPerPage < 1 {
			recordsPerPage = 10
		}

		page, err := strconv.Atoi(ginCtx.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		startIdx := (page - 1) * recordsPerPage

		matchStage := bson.D{{Key: "$match", Value: bson.D{}}}
		groupStage := bson.D{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: nil},
				{Key: "total_count", Value: bson.D{
					{Key: "$sum", Value: 1},
				}},
				{Key: "data", Value: bson.D{
					{Key: "$push", Value: "$$ROOT"},
				}},
			}},
		}
		projectStage := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "total_count", Value: 1},
				{Key: "user_items", Value: bson.D{
					{Key: "$slice", Value: []interface{}{"$data", startIdx, recordsPerPage}},
				}},
			}},
		}

		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, projectStage})
		if err != nil {
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing user items"})
			return
		}

		var allUsers []bson.M
		if err = result.All(ctx, &allUsers); err != nil {
			log.Fatal(err)
		}

		if len(allUsers) > 0 {
			ginCtx.JSON(http.StatusOK, allUsers[0])
		} else {
			ginCtx.JSON(http.StatusOK, gin.H{"message": "no users found"})
		}
	}
}

func GetUserById() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		userId := ginCtx.Param("userId")

		if err := utils.MatchUserTypeToUserId(ginCtx, userId); err != nil {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User

		err := userCollection.FindOne(ctx, bson.M{"userId": userId}).Decode(&user)
		if err != nil {
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ginCtx.JSON(http.StatusOK, user)
	}
}

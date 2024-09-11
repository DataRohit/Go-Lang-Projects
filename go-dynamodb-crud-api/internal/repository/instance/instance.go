package instance

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/datarohit/go-dynamodb-crud-api/utils/logger"
	"go.uber.org/zap"
)

func GetConnection() *dynamodb.DynamoDB {
	region := getRegion()
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String(region),
		},
	})
	if err != nil {
		logger.GetLogger().Fatal("failed to create AWS session",
			zap.Error(err),
		)
	}

	return dynamodb.New(sess)
}

func getRegion() string {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "ap-south-1"
	}
	return region
}

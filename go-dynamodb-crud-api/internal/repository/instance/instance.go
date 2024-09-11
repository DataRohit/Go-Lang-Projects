package instance

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/datarohit/go-dynamodb-crud-api/utils/logger"
	"go.uber.org/zap"
)

func GetConnection() *dynamodb.DynamoDB {
	region := getRegion()
	databaseURI := getDatabaseURI()

	config := &aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(databaseURI),
	}

	if databaseURI == "http://localhost:8000" {
		config.Credentials = credentials.NewStaticCredentials("dummy", "dummy", "")
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            *config,
	})

	if err != nil {
		logger.GetLogger().Fatal("Failed to create AWS session",
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

func getDatabaseURI() string {
	databaseURI := os.Getenv("DATABASE_URI")
	if databaseURI == "" {
		databaseURI = "http://localhost:8000"
	}
	return databaseURI
}

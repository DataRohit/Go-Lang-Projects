package instance

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetConnection() *dynamodb.DynamoDB {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "ap-south-1"
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String(region),
		},
	})
	if err != nil {
		log.Fatalf("failed to create AWS session: %v", err)
	}

	return dynamodb.New(sess)
}

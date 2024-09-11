package rules

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"io"
)

type Interface interface {
	ConvertIoReaderToStruct(reader io.Reader, model interface{}) (interface{}, error)
	GetMock() interface{}
	Migrate(connection *dynamodb.DynamoDB) error
	Validate(model interface{}) error
}

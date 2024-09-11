package product

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/datarohit/go-dynamodb-crud-api/internal/entities"
	"github.com/datarohit/go-dynamodb-crud-api/internal/entities/product"
	"github.com/datarohit/go-dynamodb-crud-api/utils/logger"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Rules struct{}

func NewRules() *Rules {
	return &Rules{}
}

func (r *Rules) ConvertIoReaderToStruct(data io.Reader, model interface{}) (interface{}, error) {
	if data == nil {
		return nil, errors.New("input reader is nil")
	}
	err := json.NewDecoder(data).Decode(model)
	if err != nil {
		logger.GetLogger().Error("Failed to decode JSON data",
			zap.Error(err),
		)
		return nil, err
	}
	return model, nil
}

func (r *Rules) Migrate(connection *dynamodb.DynamoDB) error {
	if connection == nil {
		err := errors.New("dynamodb connection is nil")
		logger.GetLogger().Error("Migration failed",
			zap.Error(err),
		)
		return err
	}

	err := r.createTable(connection)
	if err != nil {
		logger.GetLogger().Error("Table creation failed",
			zap.Error(err),
		)
		return err
	}
	return nil
}

func (r *Rules) createTable(connection *dynamodb.DynamoDB) error {
	table := &product.Product{}

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("_id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("_id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(table.TableName()),
	}

	response, err := connection.CreateTable(input)
	if err != nil {
		if strings.Contains(err.Error(), "Table already exists") {
			logger.GetLogger().Info("Table already exists, skipping creation")
			return nil
		}
		logger.GetLogger().Error("Failed to create table",
			zap.Error(err),
		)
		return err
	}

	if response != nil && strings.Contains(response.GoString(), "TableStatus: \"CREATING\"") {
		logger.GetLogger().Info("Table creation in progress, waiting...")
		time.Sleep(3 * time.Second)
		return r.createTable(connection)
	}
	return nil
}

func (r *Rules) GetMock() interface{} {
	return product.Product{
		Base: entities.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: uuid.New().String(),
	}
}

func (r *Rules) Validate(model interface{}) error {
	productModel, err := product.InterfaceToModel(model)
	if err != nil {
		logger.GetLogger().Error("Failed to convert interface to model",
			zap.Error(err),
		)
		return err
	}

	err = validation.ValidateStruct(productModel,
		validation.Field(&productModel.ID, validation.Required, is.UUIDv4),
		validation.Field(&productModel.Name, validation.Required, validation.Length(3, 50)),
	)

	if err != nil {
		logger.GetLogger().Error("Model validation failed",
			zap.Error(err),
		)
	}

	return err
}

package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/datarohit/go-dynamodb-crud-api/utils/logger"
	"go.uber.org/zap"
)

type Database struct {
	connection *dynamodb.DynamoDB
	logMode    bool
}

type Interface interface {
	Health() bool
	FindAll(condition expression.Expression, tableName string) (*dynamodb.ScanOutput, error)
	FindOne(condition map[string]interface{}, tableName string) (*dynamodb.GetItemOutput, error)
	CreateOrUpdate(entity interface{}, tableName string) (*dynamodb.PutItemOutput, error)
	Delete(condition map[string]interface{}, tableName string) (*dynamodb.DeleteItemOutput, error)
}

func NewAdapter(con *dynamodb.DynamoDB, logMode bool) Interface {
	return &Database{
		connection: con,
		logMode:    logMode,
	}
}

func (db *Database) logQuery(action string, input interface{}) {
	if db.logMode {
		logger.GetLogger().Info("DynamoDB action",
			zap.String("action", action),
			zap.Any("input", input),
		)
	}
}

func (db *Database) logError(action string, err error) {
	if db.logMode {
		logger.GetLogger().Error("DynamoDB error",
			zap.String("action", action),
			zap.Error(err),
		)
	}
}

func (db *Database) Health() bool {
	_, err := db.connection.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		db.logError("Health check", err)
		return false
	}
	return true
}

func (db *Database) FindAll(condition expression.Expression, tableName string) (*dynamodb.ScanOutput, error) {
	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  condition.Names(),
		ExpressionAttributeValues: condition.Values(),
		FilterExpression:          condition.Filter(),
		ProjectionExpression:      condition.Projection(),
		TableName:                 aws.String(tableName),
	}

	db.logQuery("FindAll", input)

	result, err := db.connection.Scan(input)
	if err != nil {
		db.logError("FindAll", err)
	}
	return result, err
}

func (db *Database) FindOne(condition map[string]interface{}, tableName string) (*dynamodb.GetItemOutput, error) {
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		db.logError("FindOne - MarshalMap", err)
		return nil, err
	}
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       conditionParsed,
	}

	db.logQuery("FindOne", input)

	result, err := db.connection.GetItem(input)
	if err != nil {
		db.logError("FindOne", err)
	}
	return result, err
}

func (db *Database) CreateOrUpdate(entity interface{}, tableName string) (*dynamodb.PutItemOutput, error) {
	entityParsed, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		db.logError("CreateOrUpdate - MarshalMap", err)
		return nil, err
	}
	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(tableName),
	}

	db.logQuery("CreateOrUpdate", input)

	result, err := db.connection.PutItem(input)
	if err != nil {
		db.logError("CreateOrUpdate", err)
	}
	return result, err
}

func (db *Database) Delete(condition map[string]interface{}, tableName string) (*dynamodb.DeleteItemOutput, error) {
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		db.logError("Delete - MarshalMap", err)
		return nil, err
	}
	input := &dynamodb.DeleteItemInput{
		Key:       conditionParsed,
		TableName: aws.String(tableName),
	}

	db.logQuery("Delete", input)

	result, err := db.connection.DeleteItem(input)
	if err != nil {
		db.logError("Delete", err)
	}
	return result, err
}

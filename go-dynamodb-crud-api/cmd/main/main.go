package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/datarohit/go-dynamodb-crud-api/config"
	"github.com/datarohit/go-dynamodb-crud-api/internal/repository/adapter"
	"github.com/datarohit/go-dynamodb-crud-api/internal/repository/instance"
	"github.com/datarohit/go-dynamodb-crud-api/internal/routes"
	"github.com/datarohit/go-dynamodb-crud-api/internal/rules"
	RulesProduct "github.com/datarohit/go-dynamodb-crud-api/internal/rules/product"
	"github.com/datarohit/go-dynamodb-crud-api/utils/logger"
	"go.uber.org/zap"
)

func main() {
	configs := config.GetConfig()
	connection := instance.GetConnection()
	repository := adapter.NewAdapter(connection, true)

	logger.GetLogger().Info("Starting service...")

	if errs := Migrate(connection); len(errs) > 0 {
		for _, err := range errs {
			logger.GetLogger().Panic("Migration error", zap.Error(err))
		}
	}

	if err := checkTables(connection); err != nil {
		logger.GetLogger().Panic("Error checking tables", zap.Error(err))
	}

	port := fmt.Sprintf(":%v", configs.Port)
	router := routes.NewRouter().SetRouters(repository)
	logger.GetLogger().Info("Service running", zap.String("port", port))

	if err := http.ListenAndServe(port, router); err != nil {
		logger.GetLogger().Fatal("Failed to start server", zap.Error(err))
	}
}

func Migrate(connection *dynamodb.DynamoDB) []error {
	var errs []error

	callMigrateAndAppendError(&errs, connection, &RulesProduct.Rules{})

	return errs
}

func callMigrateAndAppendError(errs *[]error, connection *dynamodb.DynamoDB, rule rules.Interface) {
	if err := rule.Migrate(connection); err != nil {
		*errs = append(*errs, err)
		logger.GetLogger().Error("Migration failed", zap.Error(err))
	}
}

func checkTables(connection *dynamodb.DynamoDB) error {
	response, err := connection.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		logger.GetLogger().Error("Failed to list tables", zap.Error(err))
		return err
	}

	if len(response.TableNames) == 0 {
		logger.GetLogger().Info("No tables found")
	} else {
		for _, tableName := range response.TableNames {
			logger.GetLogger().Info("Table found", zap.String("table", *tableName))
		}
	}

	return nil
}

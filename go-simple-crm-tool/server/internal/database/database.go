package database

import (
	"fmt"
	"os"
	"strconv"

	"go-simple-crm-tool/pkg/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DatabaseConnection *gorm.DB

func InitializeDatabase() error {
	databaseHost := os.Getenv("POSTGRES_HOST")
	databasePort := os.Getenv("POSTGRES_PORT")
	databaseUser := os.Getenv("POSTGRES_USER")
	databasePassword := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DB")

	port, err := strconv.Atoi(databasePort)
	if err != nil {
		utils.Fatal(logrus.Fields{
			"file": "internal/database/database.go", "action": "InitializeDatabase",
			"databasePort": databasePort,
		}, fmt.Sprintf("Invalid database port: %v", err))
		return err
	}

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		databaseHost, databaseUser, databasePassword, databaseName, port,
	)

	DatabaseConnection, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		utils.Fatal(logrus.Fields{
			"file": "internal/database/database.go", "action": "InitializeDatabase",
			"databaseHost": databaseHost,
			"databaseUser": databaseUser,
			"databaseName": databaseName,
			"databasePort": port,
		}, fmt.Sprintf("Error connecting to the database: %v", err))
		return err
	}

	utils.Info(logrus.Fields{"action": "InitializeDatabase"}, "Connected to PostgreSQL database successfully!")
	return nil
}

func CloseDatabase() error {
	sqlDatabase, err := DatabaseConnection.DB()
	if err != nil {
		utils.Warn(logrus.Fields{"action": "CloseDatabase"}, fmt.Sprintf("Failed to retrieve database object: %v", err))
		return err
	}
	if err := sqlDatabase.Close(); err != nil {
		utils.Warn(logrus.Fields{"action": "CloseDatabase"}, fmt.Sprintf("Failed to close database connection: %v", err))
		return err
	}
	utils.Info(logrus.Fields{"action": "CloseDatabase"}, "Database connection closed successfully.")
	return nil
}

func MigrateModel(model interface{}) error {
	err := DatabaseConnection.AutoMigrate(model)
	if err != nil {
		utils.Fatal(logrus.Fields{"action": "MigrateModel",
			"model": fmt.Sprintf("%T", model),
		}, fmt.Sprintf("Failed to migrate model: %v", err))
		return err
	}

	utils.Info(logrus.Fields{"action": "MigrateModel",
		"model": fmt.Sprintf("%T", model),
	}, "Model migrated successfully.")
	return nil
}

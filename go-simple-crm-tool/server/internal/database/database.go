package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

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
		log.Fatalf("Invalid database port: %v", err)
		return err
	}

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		databaseHost, databaseUser, databasePassword, databaseName, port,
	)

	DatabaseConnection, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return err
	}

	log.Println("Connected to PostgreSQL database successfully!")
	return nil
}

func CloseDatabase() error {
	sqlDatabase, err := DatabaseConnection.DB()
	if err != nil {
		return err
	}
	return sqlDatabase.Close()
}

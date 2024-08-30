package dbConfig

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DatabaseConnection *gorm.DB

func InitializeDatabase() error {
	databaseHost := "localhost"
	databasePort := 5432
	databaseUser := "pguser"
	databasePassword := "pgpass"
	databaseName := "stockdb"

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		databaseHost, databaseUser, databasePassword, databaseName, databasePort,
	)

	var err error
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

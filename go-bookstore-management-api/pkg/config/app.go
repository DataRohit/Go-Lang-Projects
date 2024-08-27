package config

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() error {
	var err error

	if err := os.MkdirAll("data", os.ModePerm); err != nil {
		log.Println("Failed to create directory:", err)
		return err
	}

	db, err = gorm.Open(sqlite.Open("data/bookstore.db"), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return err
	}

	log.Println("Connected to database successfully")
	return nil
}

func GetDB() *gorm.DB {
	return db
}

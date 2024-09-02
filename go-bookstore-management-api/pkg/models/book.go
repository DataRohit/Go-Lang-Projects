package models

import (
	"go-bookstore-management-api/pkg/config"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	Publication string    `json:"publication"`
}

func init() {
	var err error
	if err = config.Connect(); err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	db = config.GetDB()
	if db == nil {
		panic("Failed to initialize database connection")
	}

	db.AutoMigrate(&Book{})
}

func (book *Book) Create() (*Book, error) {
	book.ID = uuid.New()
	if err := db.Create(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func GetAll() ([]Book, error) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetByID(id uuid.UUID) (*Book, error) {
	var book Book
	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &book, nil
}

func DeleteByID(id uuid.UUID) (*Book, error) {
	var book Book
	if err := db.Where("id = ?", id).Delete(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &book, nil
}

func (book *Book) Update() (*Book, error) {
	if err := db.Save(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/datarohit/go-bookstore-management-api/pkg/models"
	"github.com/datarohit/go-bookstore-management-api/pkg/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

	books, err := models.GetAll()
	if err != nil {
		log.Printf("Error fetching books: %v | Method: %s | URL: %s | RemoteAddr: %s", err, r.Method, r.URL.Path, r.RemoteAddr)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching books"))
		return
	}

	res, err := json.Marshal(books)
	if err != nil {
		log.Printf("Error marshalling response: %v | Method: %s | URL: %s | RemoteAddr: %s", err, r.Method, r.URL.Path, r.RemoteAddr)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error processing response"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

	vars := mux.Vars(r)
	bookID, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("Error parsing UUID: %v | Method: %s | URL: %s | RemoteAddr: %s", err, r.Method, r.URL.Path, r.RemoteAddr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid UUID format"))
		return
	}

	book, err := models.GetByID(bookID)
	if err != nil {
		log.Printf("Error fetching book with ID %s: %v | Method: %s | URL: %s | RemoteAddr: %s", bookID, err, r.Method, r.URL.Path, r.RemoteAddr)
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Book not found"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error fetching book"))
		}
		return
	}

	res, err := json.Marshal(book)
	if err != nil {
		log.Printf("Error marshalling response: %v | Method: %s | URL: %s | RemoteAddr: %s", err, r.Method, r.URL.Path, r.RemoteAddr)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error processing response"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

	var book models.Book
	if err := utils.ParseBody(r, &book); err != nil {
		log.Printf("Error parsing request body: %v | Method: %s | URL: %s | RemoteAddr: %s", err, r.Method, r.URL.Path, r.RemoteAddr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}

	createdBook, err := book.Create()
	if err != nil {
		log.Printf("Error creating book: %v | Method: %s | URL: %s | RemoteAddr: %s", err, r.Method, r.URL.Path, r.RemoteAddr)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating book"))
		return
	}

	res, err := json.Marshal(createdBook)
	if err != nil {
		log.Printf("Error marshalling response: %v | Method: %s | URL: %s | RemoteAddr: %s", err, r.Method, r.URL.Path, r.RemoteAddr)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error processing response"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

	vars := mux.Vars(r)
	bookID, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("Error parsing UUID: %v | Method: %s | URL: %s | RemoteAddr: %s", err, r.Method, r.URL.Path, r.RemoteAddr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid UUID format"))
		return
	}

	_, err = models.DeleteByID(bookID)
	if err != nil {
		log.Printf("Error deleting book with ID %s: %v | Method: %s | URL: %s | RemoteAddr: %s", bookID, err, r.Method, r.URL.Path, r.RemoteAddr)
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Book not found"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error deleting book"))
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

	var updateBookData models.Book
	if err := utils.ParseBody(r, &updateBookData); err != nil {
		log.Printf("Error parsing request body: %v | Method: %s | URL: %s | RemoteAddr: %s", err, r.Method, r.URL.Path, r.RemoteAddr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}

	vars := mux.Vars(r)
	bookID, err := uuid.Parse(vars["id"])
	if err != nil {
		log.Printf("Error parsing UUID: %v | Method: %s | URL: %s | RemoteAddr: %s", err, r.Method, r.URL.Path, r.RemoteAddr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid UUID format"))
		return
	}

	book, err := models.GetByID(bookID)
	if err != nil {
		log.Printf("Error fetching book with ID %s: %v | Method: %s | URL: %s | RemoteAddr: %s", bookID, err, r.Method, r.URL.Path, r.RemoteAddr)
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Book not found"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error fetching book"))
		}
		return
	}

	if updateBookData.Name != "" {
		book.Name = updateBookData.Name
	}
	if updateBookData.Author != "" {
		book.Author = updateBookData.Author
	}
	if updateBookData.Publication != "" {
		book.Publication = updateBookData.Publication
	}

	updatedBook, err := book.Update()
	if err != nil {
		log.Printf("Error updating book with ID %s: %v | Method: %s | URL: %s | RemoteAddr: %s", bookID, err, r.Method, r.URL.Path, r.RemoteAddr)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error updating book"))
		return
	}

	res, err := json.Marshal(updatedBook)
	if err != nil {
		log.Printf("Error marshalling response: %v | Method: %s | URL: %s | RemoteAddr: %s", err, r.Method, r.URL.Path, r.RemoteAddr)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error processing response"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

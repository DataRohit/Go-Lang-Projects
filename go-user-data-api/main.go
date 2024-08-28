package main

import (
	"context"
	"log"
	"net/http"

	dbConfig "github.com/datarohit/go-user-data-api/config"
	"github.com/gorilla/mux"
)

func main() {
	db, err := dbConfig.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	ctx := context.TODO()
	defer func() {
		if err := db.Disconnect(ctx); err != nil {
			log.Fatalf("Error disconnecting from database: %v", err)
		}
	}()

	router := mux.NewRouter()

	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

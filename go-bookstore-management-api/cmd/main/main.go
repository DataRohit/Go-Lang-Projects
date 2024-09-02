package main

import (
	"log"
	"net/http"

	"go-bookstore-management-api/pkg/config"
	"go-bookstore-management-api/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	if err := config.Connect(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)

	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

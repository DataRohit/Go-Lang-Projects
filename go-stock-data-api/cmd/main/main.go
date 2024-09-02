package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go-stock-data-api/pkg/config/dbConfig"
	"go-stock-data-api/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	if err := dbConfig.InitializeDatabase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	defer func() {
		if err := dbConfig.CloseDatabase(); err != nil {
			log.Printf("Failed to close the database connection: %v", err)
		} else {
			log.Println("Database connection closed successfully")
		}
	}()

	router := mux.NewRouter()
	routes.RegisterStockRoutes(router)

	server := &http.Server{
		Handler:      router,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Println("Server starting on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	<-stop

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	} else {
		log.Println("Server shutdown gracefully")
	}
}

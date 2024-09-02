package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	dbConfig "go-user-data-api/config"
	userRoutes "go-user-data-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	if err := dbConfig.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer func() {
		if err := dbConfig.Disconnect(ctx); err != nil {
			log.Printf("Failed to disconnect from database: %v", err)
		} else {
			log.Println("Disconnected from database successfully")
		}
	}()

	router := mux.NewRouter()
	userRoutes.RegisterUserRoutes(router)

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

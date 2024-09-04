package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go-simple-crm-tool/api/routes"
	"go-simple-crm-tool/api/schemas"
	"go-simple-crm-tool/internal/database"
	"go-simple-crm-tool/pkg/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func closeDatabaseConnection() {
	if err := database.CloseDatabase(); err != nil {
		utils.Warn(logrus.Fields{"action": "closeDatabase"}, fmt.Sprintf("Failed to close the database connection: %v", err))
	} else {
		utils.Info(logrus.Fields{"action": "closeDatabase"}, "Database connection closed successfully")
	}
}

func startServer(router *mux.Router) *http.Server {
	server := &http.Server{
		Handler:      router,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		utils.Info(logrus.Fields{"action": "startServer"}, "Server starting on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.Warn(logrus.Fields{"action": "startServer"}, fmt.Sprintf("Server failed: %v", err))
		}
	}()

	return server
}

func waitForShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		utils.Warn(logrus.Fields{"action": "waitForShutdown"}, fmt.Sprintf("Server forced to shutdown: %v", err))
	} else {
		utils.Info(logrus.Fields{"action": "waitForShutdown"}, "Server shutdown gracefully")
	}
}

func makeMigrations() {
	err := database.MigrateModel(&schemas.Lead{})
	if err != nil {
		utils.Warn(logrus.Fields{"action": "makeMigrations"}, fmt.Sprintf("Failed to migrate Lead model: %v", err))
	}
}

func main() {
	utils.InitLogger(logrus.InfoLevel, "text", "stdout")

	if err := godotenv.Load(); err != nil {
		utils.Warn(logrus.Fields{"action": "main"}, fmt.Sprintf("Error loading .env file: %v", err))
	}

	if err := database.InitializeDatabase(); err != nil {
		utils.Warn(logrus.Fields{"action": "main"}, fmt.Sprintf("Failed to initialize database: %v", err))
	}
	defer closeDatabaseConnection()

	makeMigrations()

	router := mux.NewRouter()
	routes.RegisterLeadRoutes(router)

	server := startServer(router)
	waitForShutdown(server)
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error loading the .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/api/v1", func(ginCtx *gin.Context) {
		ginCtx.JSON(http.StatusOK, gin.H{"success": "access granted for /api/v1"})
	})

	router.GET("/api/v2", func(ginCtx *gin.Context) {
		ginCtx.JSON(http.StatusOK, gin.H{"success": "access granted for /api/v2"})
	})

	router.Run(fmt.Sprintf(":%s", port))
}

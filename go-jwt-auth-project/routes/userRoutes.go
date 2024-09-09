package routes

import (
	"github.com/datarohit/go-jwt-auth-project/handlers"
	"github.com/datarohit/go-jwt-auth-project/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.Use(middlewares.Authenticate())

	router.GET("/users", handlers.GetAllUsers())
	router.GET("/users/:userId", handlers.GetUserById())
}

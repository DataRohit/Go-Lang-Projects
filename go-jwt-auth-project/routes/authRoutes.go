package routes

import (
	"github.com/datarohit/go-jwt-auth-project/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/users/signup", handlers.SingUp())
	router.POST("/users/login", handlers.Login())
}

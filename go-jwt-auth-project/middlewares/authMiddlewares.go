package middlewares

import (
	"net/http"

	"github.com/datarohit/go-jwt-auth-project/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		clientToken := ginCtx.Request.Header.Get("token")
		if clientToken == "" {
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "no authorization header provided"})
			ginCtx.Abort()
			return
		}

		claims, err := utils.ValidateToken(clientToken)
		if err != "" {
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			ginCtx.Abort()
			return
		}

		ginCtx.Set("email", claims.Email)
		ginCtx.Set("firstName", claims.FirstName)
		ginCtx.Set("lastName", claims.LastName)
		ginCtx.Set("uid", claims.UID)
		ginCtx.Set("userType", claims.UserType)

		ginCtx.Next()
	}
}

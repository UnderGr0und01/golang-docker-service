package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
		//
		//
		if !isAuthorized(c) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func isAuthorized(c *gin.Context) bool {
	//
	//
	//
	authorizationHeader := c.GetHeader("Authorization")
	return authorizationHeader == "my-secret-token"
}

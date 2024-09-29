package middleware

import (
	"net/http"
	"next-gen-job-hunting/api/token"
	"next-gen-job-hunting/api/user"
	"next-gen-job-hunting/common/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenService *token.TokenRepository, userService *user.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("auth_token")
		if authToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "auth_token is required"})
			c.Abort()
			return
		}

		// Validate token and retrieve user
		user, err := utils.ValidateToken(authToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid auth_token"})
			c.Abort()
			return
		}

		// Store user in context
		c.Set("user", user)
		c.Next()
	}
}

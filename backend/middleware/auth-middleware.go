package middleware

import (
	"fmt"
	"net/http"
	"next-gen-job-hunting/api/token"
	"next-gen-job-hunting/api/user"
	"next-gen-job-hunting/common/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.GetHeader("auth_token"))
		c.Next()
	}
}

func AuthMiddleware(tokenService *token.TokenService, userService *user.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("auth_token")

		if authToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "auth_token is required"})
			c.Abort()
			return
		}
		if !IsValidAuthToken(authToken) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid auth_token"})
			c.Abort()
			return
		}

		tokenData, err := tokenService.FindByTokenHash(utils.GenerateTokenHash(authToken))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid auth_token"})
			c.Abort()
			return
		}

		// Store user in context
		user, err := userService.GetUserByID(tokenData.UserId, c)
		c.Set("user", user)
		fmt.Println("**************USER ID**************")
		fmt.Println(tokenData.User)
		c.Next()
	}
}

func IsValidAuthToken(authToken string) bool {
	claims, err := utils.ValidateToken(authToken)
	if err != nil {
		return false
	}
	return claims.ExpiresAt > time.Now().Unix()
}

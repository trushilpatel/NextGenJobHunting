package middleware

import (
	"fmt"
	"net/http"
	"next-gen-job-hunting/api/token"
	"next-gen-job-hunting/api/user"
	"next-gen-job-hunting/common/utils"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Logging Middleware
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.GetHeader("auth_token"))
		c.Next()
	}
}

// Authentication Middleware
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
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
			c.Abort()
			return
		}
		c.Set("user", user)
		fmt.Println("**************USER ID**************")
		fmt.Println(tokenData.User)
		c.Next()
	}
}

// CORS Middleware to allow requests from any origin
func CORSMiddleware() gin.HandlerFunc {
	config := cors.Config{
		AllowAllOrigins:  true, // Allows all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "auth_token"},
		AllowCredentials: true, // Allow credentials (e.g., cookies or auth headers)
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour, // Cache preflight requests for 12 hours
	}
	return cors.New(config)
}

// Function to validate the auth token
func IsValidAuthToken(authToken string) bool {
	claims, err := utils.ValidateToken(authToken)
	if err != nil {
		return false
	}
	return claims.ExpiresAt > time.Now().Unix()
}

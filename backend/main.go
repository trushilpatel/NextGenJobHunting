package main

import (
	"next-gen-job-hunting/api/auth"
	"next-gen-job-hunting/api/joburl"
	"next-gen-job-hunting/api/user"
	"next-gen-job-hunting/common/utils"
	"next-gen-job-hunting/config/env"
	"next-gen-job-hunting/di"
	"next-gen-job-hunting/middleware"

	"github.com/gin-gonic/gin"
)

func v1API(router *gin.Engine) {
	// Initialize necessary services
	authMiddleware := middleware.AuthMiddleware(di.InitialiseTokenService(), di.InitialiseUserService())

	// Public routes (e.g., authentication routes)
	authGroup := router.Group("/api/v1")
	auth.RegisterAuthRoutes(authGroup, di.InitializeAuthController())

	// Protected routes (requires authentication)
	v1 := router.Group("/api/v1", authMiddleware)
	{
		user.RegisterUserRoutes(v1, di.InitializeUserController())
		joburl.RegisterRoutes(v1, di.InitializeJobUrlController())
	}
}

func main() {
	// Load environment variables
	env.LoadEnvVars()

	// Run auto database migrations
	RunAutoDBMigrations()

	// Initialize Gin router
	router := gin.Default()

	// Apply CORS middleware to handle cross-origin requests from React app
	router.Use(middleware.CORSMiddleware())

	// Apply logging middleware (if any custom logging utility exists)
	router.Use(utils.Logger())

	// Register API routes (Version 1 API)
	v1API(router)

	// Start the server
	router.Run("localhost:" + env.GetPort())
}

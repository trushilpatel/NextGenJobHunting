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
	// Version 1 API
	authGroup := router.Group("/api/v1")
	auth.RegisterAuthRoutes(authGroup, di.InitializeAuthController())

	authMiddleware := middleware.AuthMiddleware(di.InitialiseTokenService())
	v1 := router.Group("/api/v1", authMiddleware)
	{
		user.RegisterUserRoutes(v1, di.InitializeUserController())
		joburl.RegisterRoutes(v1, di.InitializeJobUrlController())
	}
}

func main() {
	env.LoadEnvVars()
	RunAutoDBMigrations()

	router := gin.Default()
	router.Use(utils.Logger())
	v1API(router)
	router.Run("localhost:" + env.GetPort())
}

package main

import (
	"next-gen-job-hunting/api/auth"
	"next-gen-job-hunting/api/jobpost"
	"next-gen-job-hunting/api/joburl"
	"next-gen-job-hunting/api/user"
	"next-gen-job-hunting/common/utils"
	"next-gen-job-hunting/config/env"
	"next-gen-job-hunting/di"
	"next-gen-job-hunting/middleware"

	_ "next-gen-job-hunting/docs" // This line is necessary for go-swagger to find your docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Next Gen Job Hunting API
// @version 1.0
// @description Automating resume and cover letter generation, allowing users to easily tailor their applications to specific job opportunities while highlighting their unique skills and experiences.

// @host localhost:8080
// @BasePath /api/v1
func v1API(router *gin.Engine) {
	openRoutes := router.Group("/api/v1")

	// Swagger route
	openRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Auth routes
	auth.RegisterAuthRoutes(openRoutes, di.InitializeAuthController())
	authMiddleware := middleware.AuthMiddleware(di.InitialiseTokenService(), di.InitialiseUserService())

	v1 := router.Group("/api/v1", authMiddleware)
	{
		user.RegisterUserRoutes(v1, di.InitializeUserController())
		joburl.RegisterRoutes(v1, di.InitializeJobUrlController())
		jobpost.RegisterJobPostRoutes(v1, di.InitialiseJobPostController())
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

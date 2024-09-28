package main

import (
	"next-gen-job-hunting/api/joburl"
	"next-gen-job-hunting/api/user"
	"next-gen-job-hunting/common/utils"
	"next-gen-job-hunting/config/env"
	"next-gen-job-hunting/di"

	"github.com/gin-gonic/gin"
)

func v1API(router *gin.Engine) {
	// Version 1 API
	v1 := router.Group("/api/v1")
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
	router.Run("localhost:8080")
}

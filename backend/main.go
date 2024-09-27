package main

import (
	"net/http"
	db "next-gen-job-hunting/config/database"
	"next-gen-job-hunting/config/env"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func getCreditScore(c *gin.Context) {
	value := "No route found"
	c.IndentedJSON(http.StatusOK, value)
}

func main() {
	env.LoadEnvVars()

	db.RunAutoDBMigrations()

	router := gin.Default()
	router.GET("/", getCreditScore)
	router.Run("localhost:8080")
}

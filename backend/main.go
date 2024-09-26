package main

import (
	"net/http"
	db "next_gen_job_hunting/config/database"
	env "next_gen_job_hunting/config/env"

	"github.com/gin-gonic/gin"
)

func getCreditScore(c *gin.Context) {
	value := "No route found"
	c.IndentedJSON(http.StatusOK, value)
}

func main() {
	env.LoadEnvVars()

	db.ConnectDB()

   router := gin.Default()
   router.GET("/", getCreditScore)
   router.Run("localhost:8080")
}
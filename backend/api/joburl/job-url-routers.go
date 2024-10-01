package joburl

import "github.com/gin-gonic/gin"

// RegisterRoutes sets up the routes for Job URL operations
func RegisterRoutes(router *gin.RouterGroup, jobUrlController *JobUrlController) {
	jobRoutes := router.Group("/joburl")
	{
		jobRoutes.POST("", jobUrlController.CreateJobUrl)
		jobRoutes.GET("", jobUrlController.GetAllJobUrls)
		jobRoutes.GET("/:id", jobUrlController.GetJobUrl)
		jobRoutes.PUT("/:id", jobUrlController.UpdateJobUrl)
		jobRoutes.DELETE("/:id", jobUrlController.DeleteJobUrl)
	}
}

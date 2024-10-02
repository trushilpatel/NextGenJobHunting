package jobpost

import "github.com/gin-gonic/gin"

func RegisterJobPostRoutes(router *gin.RouterGroup, jobPostController *JobPostController) {
	jobRoutes := router.Group("/jobpost")
	{
		jobRoutes.POST("/status", jobPostController.UpdateJobPostStatus)
		jobRoutes.GET("/search", jobPostController.Search)
		jobRoutes.POST("", jobPostController.CreateJobPost)
		jobRoutes.GET("", jobPostController.GetAllJobPosts)
		jobRoutes.GET("/:id", jobPostController.GetJobPostById)
		jobRoutes.PUT("/:id", jobPostController.UpdateJobPost)
		jobRoutes.DELETE("/:id", jobPostController.DeleteJobPost)
	}
}

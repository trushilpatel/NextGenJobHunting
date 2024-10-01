package jobpost

import "github.com/gin-gonic/gin"

func RegisterJobPostRoutes(router *gin.RouterGroup, jobPostController *JobPostController) {
	jobRoutes := router.Group("/jobpost")
	{
		jobRoutes.POST("", jobPostController.CreateJobPost)
		jobRoutes.GET("", jobPostController.GetAllJobPosts)
		jobRoutes.GET("/search", jobPostController.Search)
		jobRoutes.GET("/:id", jobPostController.GetJobPostById)
		jobRoutes.PUT("/:id", jobPostController.UpdateJobPost)
		jobRoutes.DELETE("/:id", jobPostController.DeleteJobPost)
	}
}

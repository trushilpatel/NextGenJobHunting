package jobpost

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JobPostController struct {
	Service *JobPostValidationService
}

func NewJobPostController(service *JobPostValidationService) *JobPostController {
	return &JobPostController{Service: service}
}

func (controller *JobPostController) CreateJobPost(c *gin.Context) {
	var jobPost JobPost
	if err := c.ShouldBindJSON(&jobPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := controller.Service.Create(&jobPost, c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, jobPost)
}

func (controller *JobPostController) GetJobPostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	jobPost, err := controller.Service.FindByID(uint(id), c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobPost)
}

func (controller *JobPostController) Search(c *gin.Context) {
	var query JobPostQuery
	// Bind query params to the JobPostQuery struct
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobPosts, err := controller.Service.Search(query, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobPosts)
}

func (controller *JobPostController) GetAllJobPosts(c *gin.Context) {
	jobPosts, err := controller.Service.FindAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobPosts)
}

func (controller *JobPostController) UpdateJobPost(c *gin.Context) {
	var jobPost JobPost
	if err := c.ShouldBindJSON(&jobPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := controller.Service.Update(&jobPost, c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobPost)
}

func (controller *JobPostController) DeleteJobPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := controller.Service.Delete(uint(id), c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job post deleted successfully"})
}

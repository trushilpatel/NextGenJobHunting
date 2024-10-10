package jobpost

import (
	"net/http"
	"strconv"

	"next-gen-job-hunting/common/exception"

	"github.com/gin-gonic/gin"
)

// JobPostController handles job post related requests
type JobPostController struct {
	Service *JobPostValidationService
}

// NewJobPostController creates a new JobPostController
func NewJobPostController(service *JobPostValidationService) *JobPostController {
	return &JobPostController{Service: service}
}

// CreateJobPost godoc
// @Summary Create a job post
// @Description Create a new job post
// @Tags jobpost
// @Accept json
// @Produce json
// @Param jobPost body JobPost true "Job Post"
// @Success 201 {object} JobPost
// @Failure 400 {object} exception.CommonException "Invalid request payload"
// @Failure 500 {object} exception.CommonException "Failed to create job post"
// @Router /jobposts [post]
func (controller *JobPostController) CreateJobPost(c *gin.Context) {
	var jobPost JobPost
	if err := c.ShouldBindJSON(&jobPost); err != nil {
		c.JSON(http.StatusBadRequest, exception.NewCommonException("Invalid request payload", err.Error()))
		return
	}
	if err := controller.Service.Create(&jobPost, c); err != nil {
		c.JSON(http.StatusInternalServerError, exception.NewCommonException("Failed to create job post", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, jobPost)
}

// GetJobPostById godoc
// @Summary Get a job post by ID
// @Description Get a job post by its ID
// @Tags jobpost
// @Produce json
// @Param id path int true "Job Post ID"
// @Success 200 {object} JobPost
// @Failure 400 {object} exception.CommonException "Invalid ID"
// @Failure 404 {object} exception.CommonException "Job post not found"
// @Router /jobposts/{id} [get]
func (controller *JobPostController) GetJobPostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, exception.NewCommonException("Invalid ID", err.Error()))
		return
	}

	jobPost, err := controller.Service.FindByID(uint(id), c)
	if err != nil {
		c.JSON(http.StatusNotFound, exception.NewCommonException("Job post not found", err.Error()))
		return
	}
	c.JSON(http.StatusOK, jobPost)
}

// Search godoc
// @Summary Search job posts
// @Description Search job posts based on query parameters
// @Tags jobpost
// @Produce json
// @Param query query JobPostQuery true "Job Post Query"
// @Success 200 {object} common.PaginationData{data=[]JobPost}
// @Failure 400 {object} exception.CommonException "Invalid query parameters"
// @Failure 500 {object} exception.CommonException "Failed to search job posts"
// @Router /jobposts/search [get]
func (controller *JobPostController) Search(c *gin.Context) {
	var query JobPostQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, exception.NewCommonException("Invalid query parameters", err.Error()))
		return
	}

	data, err := controller.Service.Search(query, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, exception.NewCommonException("Failed to search job posts", err.Error()))
		return
	}
	c.JSON(http.StatusOK, data)
}

// GetAllJobPosts godoc
// @Summary Get all job posts
// @Description Get all job posts
// @Tags jobpost
// @Produce json
// @Success 200 {array} JobPost
// @Failure 500 {object} exception.CommonException "Failed to retrieve job posts"
// @Router /jobposts [get]
func (controller *JobPostController) GetAllJobPosts(c *gin.Context) {
	jobPosts, err := controller.Service.FindAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, exception.NewCommonException("Failed to retrieve job posts", err.Error()))
		return
	}
	c.JSON(http.StatusOK, jobPosts)
}

// UpdateJobPost godoc
// @Summary Update a job post
// @Description Update an existing job post
// @Tags jobpost
// @Accept json
// @Produce json
// @Param jobPost body JobPost true "Job Post"
// @Success 200 {object} JobPost
// @Failure 400 {object} exception.CommonException "Invalid request payload"
// @Failure 500 {object} exception.CommonException "Failed to update job post"
// @Router /jobposts [put]
func (controller *JobPostController) UpdateJobPost(c *gin.Context) {
	var jobPost JobPost
	if err := c.ShouldBindJSON(&jobPost); err != nil {
		c.JSON(http.StatusBadRequest, exception.NewCommonException("Invalid request payload", err.Error()))
		return
	}
	if err := controller.Service.Update(&jobPost, c); err != nil {
		c.JSON(http.StatusInternalServerError, exception.NewCommonException("Failed to update job post", err.Error()))
		return
	}
	c.JSON(http.StatusOK, jobPost)
}

// UpdateJobPostStatus godoc
// @Summary Update job post status
// @Description Update the status of a job post
// @Tags jobpost
// @Accept json
// @Produce json
// @Param updateUserJobPostDto body JobPostUserJobPostDto true "Update Job Post Status"
// @Success 200 {object} JobPostUserJobPostDto
// @Failure 400 {object} exception.CommonException "Invalid request payload"
// @Failure 500 {object} exception.CommonException "Failed to update job post status"
// @Router /jobposts/status [put]
func (controller *JobPostController) UpdateJobPostStatus(c *gin.Context) {
	var updateUserJobPostDto JobPostUserJobPostDto
	if err := c.ShouldBindJSON(&updateUserJobPostDto); err != nil {
		c.JSON(http.StatusBadRequest, exception.NewCommonException("Invalid request payload", err.Error()))
		return
	}

	if _, err := controller.Service.UpdateJobPostStatus(&updateUserJobPostDto, c); err != nil {
		c.JSON(http.StatusInternalServerError, exception.NewCommonException("Failed to update job post status", err.Error()))
		return
	}
	c.JSON(http.StatusOK, updateUserJobPostDto)
}

// DeleteJobPost godoc
// @Summary Delete a job post
// @Description Delete a job post by ID
// @Tags jobpost
// @Produce json
// @Param id path int true "Job Post ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} exception.CommonException "Invalid ID"
// @Failure 500 {object} exception.CommonException "Failed to delete job post"
// @Router /jobposts/{id} [delete]
func (controller *JobPostController) DeleteJobPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, exception.NewCommonException("Invalid ID", err.Error()))
		return
	}
	if err := controller.Service.Delete(uint(id), c); err != nil {
		c.JSON(http.StatusInternalServerError, exception.NewCommonException("Failed to delete job post", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job post deleted successfully"})
}

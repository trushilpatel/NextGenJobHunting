package joburl

import (
	"net/http"
	"strconv"

	"next-gen-job-hunting/api/user"

	"github.com/gin-gonic/gin"
)

type JobUrlController struct {
	Service     *JobUrlService
	UserService *user.UserService
}

func NewJobUrlController(service *JobUrlService, userService *user.UserService) *JobUrlController {
	return &JobUrlController{Service: service, UserService: userService}
}

func (controller *JobUrlController) CreateJobUrl(c *gin.Context) {
	var job JobUrl
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := controller.UserService.GetUserByID(job.UserID, c); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := controller.Service.CreateJobUrl(&job, c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, job)
}

func (controller *JobUrlController) GetJobUrl(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	jobUrl, err := controller.Service.GetJobUrlById(uint(id), c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job URL not found"})
		return
	}

	c.JSON(http.StatusOK, jobUrl)
}

func (controller *JobUrlController) GetAllJobUrls(c *gin.Context) {
	var jobs []*JobUrl
	jobs, error := controller.Service.GetAllJobUrl(c)
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}
	c.JSON(http.StatusOK, jobs)
}

func (controller *JobUrlController) UpdateJobUrl(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var job JobUrl
	if _, err := controller.Service.GetJobUrlById(uint(id), c); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job URL not found"})
		return
	}

	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := controller.UserService.GetUserByID(uint(job.UserID), c); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	job.ID.ID = uint(id)
	if err := controller.Service.UpdateJobUrl(&job, c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, job)
}

func (controller *JobUrlController) DeleteJobUrl(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if _, err := controller.Service.GetJobUrlById(uint(id), c); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job URL not found"})
		return
	}

	if err := controller.Service.DeleteJobUrl(uint(id), c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job URL deleted successfully"})
}

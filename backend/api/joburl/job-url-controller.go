package joburl

import (
	"net/http"
	"strconv"

	"next-gen-job-hunting/api/user"
	"next-gen-job-hunting/common/exception"

	"github.com/gin-gonic/gin"
)

type JobUrlController struct {
	Service     *JobUrlService
	UserService *user.UserService
}

func NewJobUrlController(service *JobUrlService, userService *user.UserService) *JobUrlController {
	return &JobUrlController{Service: service, UserService: userService}
}

// CreateJobUrl handles the creation of a new job URL.
// @Summary Create a new job URL
// @Description Create a new job URL for a user
// @Tags joburl
// @Accept json
// @Produce json
// @Param job body JobUrl true "Job URL"
// @Success 201 {object} JobUrl
// @Failure 400 {object} exception.CommonException "Invalid request payload"
// @Failure 404 {object} exception.CommonException "User not found"
// @Failure 500 {object} exception.CommonException "Could not create Job URL"
// @Router /joburl [post]
func (controller *JobUrlController) CreateJobUrl(c *gin.Context) {
	var job JobUrl
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, exception.NewCommonException("Invalid request payload", err.Error()))
		return
	}

	if _, err := controller.UserService.GetUserByID(job.UserID, c); err != nil {
		c.JSON(http.StatusNotFound, exception.NewCommonException("User not found", err.Error()))
		return
	}

	if err := controller.Service.CreateJobUrl(&job, c); err != nil {
		c.JSON(http.StatusInternalServerError, exception.NewCommonException("Could not create Job URL", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, job)
}

// GetJobUrl handles fetching a job URL by ID.
// @Summary Get a job URL by ID
// @Description Get a job URL by its ID
// @Tags joburl
// @Produce json
// @Param id path int true "Job URL ID"
// @Success 200 {object} JobUrl
// @Failure 400 {object} exception.CommonException "Invalid user ID"
// @Failure 404 {object} exception.CommonException "Job URL not found"
// @Failure 500 {object} exception.CommonException "Could not fetch Job URL"
// @Router /joburl/{id} [get]
func (controller *JobUrlController) GetJobUrl(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, exception.NewCommonException("Invalid user ID", err.Error()))
		return
	}

	jobUrl, err := controller.Service.GetJobUrlById(uint(id), c)
	if err != nil {
		c.JSON(http.StatusNotFound, exception.NewCommonException("Job URL not found", err.Error()))
		return
	}

	c.JSON(http.StatusOK, jobUrl)
}

// GetAllJobUrls handles fetching all job URLs.
// @Summary Get all job URLs
// @Description Get all job URLs
// @Tags joburl
// @Produce json
// @Success 200 {array} JobUrl
// @Failure 500 {object} exception.CommonException "Could not fetch job URLs"
// @Router /joburl [get]
func (controller *JobUrlController) GetAllJobUrls(c *gin.Context) {
	jobs, err := controller.Service.GetAllJobUrl(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, exception.NewCommonException("Could not fetch job URLs", err.Error()))
		return
	}
	c.JSON(http.StatusOK, jobs)
}

// UpdateJobUrl handles updating a job URL by ID.
// @Summary Update a job URL by ID
// @Description Update a job URL by its ID
// @Tags joburl
// @Accept json
// @Produce json
// @Param id path int true "Job URL ID"
// @Param job body JobUrl true "Job URL"
// @Success 200 {object} JobUrl
// @Failure 400 {object} exception.CommonException "Invalid user ID"
// @Failure 404 {object} exception.CommonException "Job URL not found"
// @Failure 404 {object} exception.CommonException "User not found"
// @Failure 500 {object} exception.CommonException "Could not update Job URL"
// @Router /joburl/{id} [put]
func (controller *JobUrlController) UpdateJobUrl(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, exception.NewCommonException("Invalid user ID", err.Error()))
		return
	}

	var job JobUrl
	if _, err := controller.Service.GetJobUrlById(uint(id), c); err != nil {
		c.JSON(http.StatusNotFound, exception.NewCommonException("Job URL not found", err.Error()))
		return
	}

	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, exception.NewCommonException("Invalid request payload", err.Error()))
		return
	}

	if _, err := controller.UserService.GetUserByID(uint(job.UserID), c); err != nil {
		c.JSON(http.StatusNotFound, exception.NewCommonException("User not found", err.Error()))
		return
	}

	job.ID.ID = uint(id)
	if err := controller.Service.UpdateJobUrl(&job, c); err != nil {
		c.JSON(http.StatusInternalServerError, exception.NewCommonException("Could not update Job URL", err.Error()))
		return
	}

	c.JSON(http.StatusOK, job)
}

// DeleteJobUrl handles deleting a job URL by ID.
// @Summary Delete a job URL by ID
// @Description Delete a job URL by its ID
// @Tags joburl
// @Produce json
// @Param id path int true "Job URL ID"
// @Success 200 {object} map[string]string "message: Job URL deleted successfully"
// @Failure 400 {object} exception.CommonException "Invalid user ID"
// @Failure 404 {object} exception.CommonException "Job URL not found"
// @Failure 500 {object} exception.CommonException "Could not delete Job URL"
// @Router /joburl/{id} [delete]
func (controller *JobUrlController) DeleteJobUrl(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, exception.NewCommonException("Invalid user ID", err.Error()))
		return
	}

	if _, err := controller.Service.GetJobUrlById(uint(id), c); err != nil {
		c.JSON(http.StatusNotFound, exception.NewCommonException("Job URL not found", err.Error()))
		return
	}

	if err := controller.Service.DeleteJobUrl(uint(id), c); err != nil {
		c.JSON(http.StatusInternalServerError, exception.NewCommonException("Could not delete Job URL", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job URL deleted successfully"})
}

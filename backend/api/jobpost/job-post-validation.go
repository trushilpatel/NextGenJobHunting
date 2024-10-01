package jobpost

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type JobPostValidationService struct {
	Service *JobPostService
}

func NewJobPostValidationService(service *JobPostService) *JobPostValidationService {
	return &JobPostValidationService{Service: service}
}

func (s *JobPostValidationService) Create(jobPost *JobPost, c *gin.Context) error {
	if err := validateJobPost(jobPost); err != nil {
		return err
	}
	return s.Service.Create(jobPost, c)
}

func (s *JobPostValidationService) Search(query JobPostQuery, c *gin.Context) ([]JobPost, error) {
	query.Validate()
	return s.Service.Search(query, c)
}

func (s *JobPostValidationService) FindAll(c *gin.Context) ([]JobPost, error) {
	return s.Service.FindAll(c)
}

func (s *JobPostValidationService) FindByID(id uint, c *gin.Context) (*JobPost, error) {
	if id == 0 {
		return nil, errors.New("invalid ID")
	}
	return s.Service.FindByID(id, c)
}

func (s *JobPostValidationService) Update(jobPost *JobPost, c *gin.Context) error {
	if err := validateJobPost(jobPost); err != nil {
		return err
	}
	return s.Service.Update(jobPost, c)
}

func (s *JobPostValidationService) Delete(id uint, c *gin.Context) error {
	if id == 0 {
		return errors.New("invalid job post ID")
	}
	return s.Service.Delete(id, c)
}

func validateJobPost(jobPost *JobPost) error {
	if jobPost.ID.ID != 0 {
		return errors.New("Id should not be set")
	}
	if jobPost.JobTitle == "" {
		return errors.New("jobTitle is required")
	}
	if jobPost.JobID == "" {
		return errors.New("jobId title is required")
	}
	if jobPost.ApplicationLink == "" {
		return errors.New("ApplicationLink is required")
	}
	if jobPost.JobDescription == "" {
		return errors.New("job post description is required")
	}
	if jobPost.JobPostedDate.IsZero() {
		return errors.New("job post date is required")
	}
	return nil
}

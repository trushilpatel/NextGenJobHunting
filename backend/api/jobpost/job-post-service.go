package jobpost

import (
	"github.com/gin-gonic/gin"
)

type JobPostService struct {
	Repo *JobPostRepository
}

func NewJobPostService(repo *JobPostRepository) *JobPostService {
	return &JobPostService{Repo: repo}
}

func (s *JobPostService) Create(jobPost *JobPost, c *gin.Context) error {
	return s.Repo.Create(jobPost, c)
}

func (s *JobPostService) FindAll(c *gin.Context) ([]JobPost, error) {
	return s.Repo.FindAll(c)
}

func (s *JobPostService) FindByID(id uint, c *gin.Context) (*JobPost, error) {
	return s.Repo.FindByID(id, c)
}

func (s *JobPostService) Update(jobPost *JobPost, c *gin.Context) error {
	return s.Repo.Update(jobPost, c)
}

func (s *JobPostService) Delete(id uint, c *gin.Context) error {
	return s.Repo.Delete(id, c)
}

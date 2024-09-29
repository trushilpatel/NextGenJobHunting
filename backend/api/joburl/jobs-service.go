package joburl

import "github.com/gin-gonic/gin"

type JobUrlService struct {
	repo *JobUrlRepository
}

func NewJobUrlService(repo *JobUrlRepository) *JobUrlService {
	return &JobUrlService{repo: repo}
}

func (s *JobUrlService) CreateJobUrl(jobUrl *JobUrl, c *gin.Context) error {
	return s.repo.CreateJob(jobUrl, c)
}

func (s *JobUrlService) GetAllJobUrl(c *gin.Context) ([]*JobUrl, error) {
	return s.repo.GetAllJobUrl(c)
}

func (s *JobUrlService) GetJobUrlById(id uint, c *gin.Context) (*JobUrl, error) {
	return s.repo.GetJobUrlById(id, c)
}

func (s *JobUrlService) UpdateJobUrl(jobUrl *JobUrl, c *gin.Context) error {
	return s.repo.UpdateJobUrl(jobUrl, c)
}

func (s *JobUrlService) DeleteJobUrl(id uint, c *gin.Context) error {
	return s.repo.DeleteJobUrl(id, c)
}

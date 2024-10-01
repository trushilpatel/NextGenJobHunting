package jobpost

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JobPostRepository struct {
	DB *gorm.DB
}

func NewJobPostRepository(db *gorm.DB) *JobPostRepository {
	return &JobPostRepository{DB: db}
}

func (r *JobPostRepository) Create(jobPost *JobPost, c *gin.Context) error {
	return r.DB.Create(jobPost).Error
}

func (r *JobPostRepository) FindAll(c *gin.Context) ([]JobPost, error) {
	var jobPosts []JobPost
	err := r.DB.Find(&jobPosts).Error
	return jobPosts, err
}

func (r *JobPostRepository) FindByID(id uint, c *gin.Context) (*JobPost, error) {
	var jobPost JobPost
	err := r.DB.First(&jobPost, id).Error
	return &jobPost, err
}

func (r *JobPostRepository) Update(jobPost *JobPost, c *gin.Context) error {
	return r.DB.Save(jobPost).Error
}

func (r *JobPostRepository) Delete(id uint, c *gin.Context) error {
	return r.DB.Delete(&JobPost{}, id).Error
}

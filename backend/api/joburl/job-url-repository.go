package joburl

import (
	"next-gen-job-hunting/api/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JobUrlRepository struct {
	DB *gorm.DB
}

func NewJobUrlRepository(db *gorm.DB) *JobUrlRepository {
	return &JobUrlRepository{DB: db}
}

func (r *JobUrlRepository) CreateJob(jobUrl *JobUrl, c *gin.Context) error {
	u, _ := c.Get("user")

	jobUrl.UserID = u.(*user.User).ID.ID

	if err := r.DB.Create(jobUrl).Error; err != nil {
		return err
	}
	return nil
}

func (r *JobUrlRepository) GetAllJobUrl(c *gin.Context) ([]*JobUrl, error) {
	u, _ := c.Get("user")

	var jobUrls []*JobUrl
	if err := r.DB.Where("user_id = ?", u.(*user.User).ID.ID).Find(&jobUrls).Error; err != nil {
		return nil, err
	}
	return jobUrls, nil
}

func (r *JobUrlRepository) GetJobUrlById(id uint, c *gin.Context) (*JobUrl, error) {
	u, _ := c.Get("user")

	var jobUrl JobUrl
	if err := r.DB.Where("id = ? AND user_id = ?", id, u.(*user.User).ID.ID).First(&jobUrl).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &jobUrl, nil
}

func (r *JobUrlRepository) UpdateJobUrl(jobUrl *JobUrl, c *gin.Context) error {
	u, _ := c.Get("user")

	if err := r.DB.Where("id = ? AND user_id = ?", jobUrl.ID, u.(*user.User).ID.ID).Save(jobUrl).Error; err != nil {
		return err
	}
	return nil
}

func (r *JobUrlRepository) DeleteJobUrl(id uint, c *gin.Context) error {
	if err := r.DB.Delete(&JobUrl{}, id).Error; err != nil {
		return err
	}
	return nil
}

package user_job_post

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserJobPostRepository struct {
	Db *gorm.DB
}

func NewUserJobPostRepository(db *gorm.DB, c *gin.Context) *UserJobPostRepository {
	return &UserJobPostRepository{Db: db}
}

func (r *UserJobPostRepository) Create(userJobPost *UserJobPost, c *gin.Context) (*UserJobPost, error) {
	if err := r.Db.Create(userJobPost).Error; err != nil {
		return nil, errors.New("failed to create user job post")
	}
	return userJobPost, nil
}

func (r *UserJobPostRepository) Update(userJobPost *UserJobPost, c *gin.Context) (*UserJobPost, error) {
	if err := r.Db.Save(userJobPost).Error; err != nil {
		return nil, errors.New("failed to update user job post")
	}
	return userJobPost, nil
}

func (r *UserJobPostRepository) Delete(id uint, c *gin.Context) error {
	if err := r.Db.Delete(&UserJobPost{}, id).Error; err != nil {
		return errors.New("failed to delete user job post")
	}
	return nil
}

func (r *UserJobPostRepository) FindById(id uint, c *gin.Context) (*UserJobPost, error) {
	var userJobPost UserJobPost
	if err := r.Db.First(&userJobPost, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user job post not found")
		}
		return nil, errors.New("failed to find user job post")
	}
	return &userJobPost, nil
}

func (r *UserJobPostRepository) Find(c *gin.Context) ([]UserJobPost, error) {
	var userJobPosts []UserJobPost
	if err := r.Db.Find(&userJobPosts).Error; err != nil {
		return nil, errors.New("failed to find user job posts")
	}
	return userJobPosts, nil
}

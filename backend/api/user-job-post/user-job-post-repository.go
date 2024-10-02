package user_job_post

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserJobPostRepository struct {
	Db *gorm.DB
}

func NewUserJobPostRepository(db *gorm.DB) *UserJobPostRepository {
	return &UserJobPostRepository{Db: db}
}

func (r *UserJobPostRepository) Create(userJobPost *UserJobPost, c *gin.Context) (*UserJobPost, error) {
	if err := r.Db.Create(userJobPost).Error; err != nil {
		return nil, errors.New("failed to create user job post")
	}
	return userJobPost, nil
}

func (v *UserJobPostRepository) FindByJobPostIDAndUserId(userId uint, jobPostID uint, c *gin.Context) (*UserJobPost, error) {
	var userJobPost UserJobPost
	if err := v.Db.Where("user_id = ? AND job_post_id = ?", userId, jobPostID).First(&userJobPost).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user job post not found")
		}
		return nil, errors.New("failed to find user job post")
	}
	return &userJobPost, nil
}

func (r *UserJobPostRepository) Update(userJobPost *UserJobPost, c *gin.Context) (*UserJobPost, error) {
	if err := r.Db.Where("job_post_id = ? AND user_id = ?", userJobPost.JobPostId, userJobPost.UserId).Save(userJobPost).Error; err != nil {
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

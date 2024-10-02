package user_job_post

import (
	"errors"
	"regexp"

	"github.com/gin-gonic/gin"
)

type UserJobPostValidationService struct {
	Service *UserJobPostService
}

func NewUserJobPostValidationService(userJobPostService *UserJobPostService) *UserJobPostValidationService {
	return &UserJobPostValidationService{
		Service: userJobPostService,
	}
}

func (v *UserJobPostValidationService) CreateUserJobPost(userJobPost *UserJobPost, c *gin.Context) (*UserJobPost, error) {
	if err := v.validateJobPost(userJobPost); err != nil {
		return nil, err
	}
	return v.Service.CreateUserJobPost(userJobPost, c)
}

func (v *UserJobPostValidationService) FindByJobPostIDAndUserId(userId uint, jobPostID uint, c *gin.Context) (*UserJobPost, error) {
	if userId == 0 {
		return nil, errors.New("invalid UserId: must be non-empty")
	}
	if jobPostID == 0 {
		return nil, errors.New("invalid JobPostId: must be non-empty")
	}
	return v.Service.FindByJobPostIDAndUserId(userId, jobPostID, c)
}

func (v *UserJobPostValidationService) UpdateUserJobPost(userJobPost *UserJobPost, c *gin.Context) (*UserJobPost, error) {
	if err := v.validateJobPost(userJobPost); err != nil {
		return nil, err
	}
	return v.Service.UpdateUserJobPost(userJobPost, c)
}

func (v *UserJobPostValidationService) DeleteUserJobPost(jobPostID uint, c *gin.Context) error {
	return v.Service.DeleteUserJobPost(jobPostID, c)
}

func (v *UserJobPostValidationService) validateJobPost(userJobPost *UserJobPost) error {
	if userJobPost.JobPostId == 0 {
		return errors.New("invalid JobPostId: must be non-empty")
	}
	if userJobPost.UserId == 0 {
		return errors.New("invalid UserId: must be non-empty")
	}
	if userJobPost.JobApplicationStatus == "" {
		return errors.New("invalid JobApplicationStatus: must be non-empty")
	}
	if userJobPost.ResumeScore < 0 || userJobPost.ResumeScore > 100 {
		return errors.New("invalid ResumeScore: must be between 0 and 100")
	}
	return nil
}

func isValidURL(str string) bool {
	re := regexp.MustCompile(`^(http|https):\/\/[^\s$.?#].[^\s]*$`)
	return re.MatchString(str)
}

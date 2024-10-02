package jobpost

import (
	"fmt"
	"next-gen-job-hunting/api/user"
	user_job_post "next-gen-job-hunting/api/user-job-post"

	"github.com/gin-gonic/gin"
)

type JobPostService struct {
	Repo               *JobPostRepository
	UserJobPostService *user_job_post.UserJobPostValidationService
}

func NewJobPostService(repo *JobPostRepository, userJobPostService *user_job_post.UserJobPostValidationService) *JobPostService {
	return &JobPostService{Repo: repo, UserJobPostService: userJobPostService}
}

func (s *JobPostService) Create(jobPost *JobPost, c *gin.Context) error {
	return s.Repo.Create(jobPost, c)
}

func (s *JobPostService) UpdateJobPostStatus(updateJobPostDto *JobPostUserJobPostDto, c *gin.Context) (*JobPostUserJobPostDto, error) {
	u, exists := c.Get("user")
	if !exists {
		return nil, fmt.Errorf("user not found in context")
	}
	user, ok := u.(*user.User)
	if !ok {
		return nil, fmt.Errorf("user type assertion failed")
	}

	userId := user.ID.ID
	existingUserJobPost, err := s.UserJobPostService.FindByJobPostIDAndUserId(userId, updateJobPostDto.ID.ID, c)

	if err != nil && existingUserJobPost == nil {
		// UserJobPost does not exist, create a new one
		newUserJobPost := &user_job_post.UserJobPost{
			JobPostId:            updateJobPostDto.ID.ID,
			UserId:               userId,
			JobApplicationStatus: updateJobPostDto.JobApplicationStatus,
		}
		_, err = s.UserJobPostService.CreateUserJobPost(newUserJobPost, c)
		if err != nil {
			return nil, err
		}
		updateJobPostDto.UserJobPost = *newUserJobPost
		return updateJobPostDto, nil
	} else {
		// Update the existing UserJobPost
		existingUserJobPost.JobApplicationStatus = updateJobPostDto.JobApplicationStatus
		savedDto, err := s.UserJobPostService.UpdateUserJobPost(existingUserJobPost, c)
		if err != nil {
			return nil, err
		}
		updateJobPostDto.UserJobPost = *savedDto

		return updateJobPostDto, nil
	}
}

func (s *JobPostService) Search(query JobPostQuery, c *gin.Context) ([]JobPostUserJobPostDto, error) {
	return s.Repo.Search(query, c)
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

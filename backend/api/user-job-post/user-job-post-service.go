package user_job_post

import "github.com/gin-gonic/gin"

type UserJobPostService struct {
	UserJobPostRepository *UserJobPostRepository
}

func NewUserJobPostService(userJobPostRepository *UserJobPostRepository) *UserJobPostService {
	return &UserJobPostService{
		UserJobPostRepository: userJobPostRepository,
	}
}

func (s *UserJobPostService) CreateUserJobPost(userJobPost *UserJobPost, c *gin.Context) (*UserJobPost, error) {
	return s.UserJobPostRepository.Create(userJobPost, c)
}

func (s *UserJobPostService) UpdateUserJobPost(userJobPost *UserJobPost, c *gin.Context) (*UserJobPost, error) {
	return s.UserJobPostRepository.Update(userJobPost, c)
}

func (s *UserJobPostService) DeleteUserJobPost(id uint, c *gin.Context) error {
	return s.UserJobPostRepository.Delete(id, c)
}

func (s *UserJobPostService) GetUserJobPostById(id uint, c *gin.Context) (*UserJobPost, error) {
	return s.UserJobPostRepository.FindById(id, c)
}

func (s *UserJobPostService) GetAllUserJobPosts(c *gin.Context) ([]UserJobPost, error) {
	return s.UserJobPostRepository.Find(c)
}

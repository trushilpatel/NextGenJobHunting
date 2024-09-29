package user

import "github.com/gin-gonic/gin"

type UserService struct {
	Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *User, c *gin.Context) error {
	return s.Repo.CreateUser(user, c)
}

func (s *UserService) GetAllUser(c *gin.Context) ([]*User, error) {
	return s.Repo.GetAllUser(c)
}

func (s *UserService) GetUserByID(id uint, c *gin.Context) (*User, error) {
	return s.Repo.GetUserByID(id, c)
}

func (s *UserService) GetUserByEmail(email string, c *gin.Context) (*User, error) {
	return s.Repo.GetUserByEmail(email, c)
}

func (s *UserService) UpdateUser(user *User, c *gin.Context) error {
	return s.Repo.UpdateUser(user, c)
}

func (s *UserService) DeleteUser(id uint, c *gin.Context) error {
	return s.Repo.DeleteUser(id, c)
}

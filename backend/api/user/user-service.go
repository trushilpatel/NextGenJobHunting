package user

type UserService struct {
	Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *User) error {
	return s.Repo.CreateUser(user)
}

func (s *UserService) GetAllUser() ([]*User, error) {
	return s.Repo.GetAllUser()
}

func (s *UserService) GetUserByID(id uint) (*User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*User, error) {
	return s.Repo.GetUserByEmail(email)
}

func (s *UserService) UpdateUser(user *User) error {
	return s.Repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.Repo.DeleteUser(id)
}

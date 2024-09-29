package user

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidEmail    = errors.New("invalid email address")
	ErrInvalidUsername = errors.New("username must be at least 3 characters")
	RE                 = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

type UserValidationService struct {
	Service *UserService
}

func NewUserValidationService(service *UserService) *UserValidationService {
	return &UserValidationService{Service: service}
}

func (s *UserValidationService) CreateUser(user *User) error {
	if err := ValidateUser(user); err != nil {
		return err
	}
	return s.Service.CreateUser(user)
}

func (s *UserValidationService) GetAllUser() ([]*User, error) {
	return s.Service.GetAllUser()
}

func (s *UserValidationService) GetUserByID(id uint) (*User, error) {
	return s.Service.GetUserByID(id)
}

func (s *UserValidationService) UpdateUser(user *User) error {
	if err := ValidateUser(user); err != nil {
		return err
	}
	return s.Service.UpdateUser(user)
}

func (s *UserValidationService) DeleteUser(id uint) error {
	user, err := s.Service.GetUserByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	return s.Service.DeleteUser(id)
}

func ValidateUser(user *User) error {
	if err := validateEmail(user.Email); err != nil {
		return err
	}
	if err := validateUsername(user.Username); err != nil {
		return err
	}

	return nil
}

func validateEmail(email string) error {
	if !RE.MatchString(email) {
		return ErrInvalidEmail
	}
	return nil
}

func validateUsername(username string) error {
	if len(username) < 3 {
		return ErrInvalidUsername
	}
	return nil
}

package auth

import (
	"fmt"
	"next-gen-job-hunting/api/token"
	"next-gen-job-hunting/api/user"
	"next-gen-job-hunting/common/utils"
)

type AuthValidationService struct {
	Service *AuthService
}

func NewAuthValidator(service *AuthService) *AuthValidationService {
	return &AuthValidationService{
		Service: service,
	}
}

func (s *AuthValidationService) SignUp(user *user.User) (*token.Token, error) {
	_, err := s.ValidateSignUp(user)
	if err != nil {
		return nil, err
	}

	token, err := s.Service.SignUp(user)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *AuthValidationService) SignIn(user *user.User) (*token.Token, error) {
	_, err := s.ValidateSignIn(user)
	if err != nil {
		return nil, err
	}

	token, err := s.Service.SignIn(user)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *AuthValidationService) SignOut(token string) error {
	_, err := utils.ValidateToken(token)
	if err != nil {
		return err
	}

	err = s.Service.SignOut(token)
	if err != nil {
		return err
	}
	return nil
}

func (v *AuthValidationService) AuthenticateUser(user *user.User) (bool, error) {
	authenticatedUser, err := v.Service.GetUserByEmail(user.Email)
	if err != nil || authenticatedUser == nil {
		return false, err
	}

	// To Do implement encryption here and in user-service.go
	// to encrypt the password saving
	// create new API for change password
	if authenticatedUser.Password != user.Password {
		return false, nil
	}

	return true, nil
}

func (v *AuthValidationService) ValidateSignUp(user *user.User) (*user.User, error) {
	existingUser, err := v.Service.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, fmt.Errorf("email already in use")
	}

	return user, nil
}

func (v *AuthValidationService) ValidateSignIn(user *user.User) (*user.User, error) {
	isAuthenticatedUser, err := v.AuthenticateUser(user)
	if err != nil || !isAuthenticatedUser {
		return nil, err
	}

	return user, nil
}

func NewAuthValidationService(service *AuthService) *AuthValidationService {
	return &AuthValidationService{
		Service: service,
	}
}

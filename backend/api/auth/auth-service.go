package auth

import (
	"next-gen-job-hunting/api/user"
	"next-gen-job-hunting/common/utils"

	"next-gen-job-hunting/api/token"
)

type AuthService struct {
	UserRepository *user.UserRepository
	TokenService   *token.TokenService
}

func NewAuthService(UserRepository *user.UserRepository, tokenService *token.TokenService) *AuthService {
	return &AuthService{
		UserRepository: UserRepository,
		TokenService:   tokenService,
	}
}

func (s *AuthService) SignUp(user *user.User) (*token.Token, error) {
	if err := s.UserRepository.CreateUser(user); err != nil {
		return nil, err
	}

	token, err := s.TokenService.CreateTokenForUser(user)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *AuthService) SignIn(user *user.User) (*token.Token, error) {
	token, err := s.TokenService.CreateTokenForUser(user)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *AuthService) SignOut(token string) error {
	_, err := s.TokenService.DeleteTokenByTokenHash(utils.GenerateTokenHash(token))
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) GetUserByEmail(email string) (*user.User, error) {
	return s.UserRepository.GetUserByEmail(email)
}

func (s *AuthService) AuthenticateUser(user *user.User) (bool, error) {
	authenticatedUser, err := s.GetUserByEmail(user.Email)
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

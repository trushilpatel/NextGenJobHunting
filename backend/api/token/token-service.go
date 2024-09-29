package token

import (
	"next-gen-job-hunting/api/user"
	"next-gen-job-hunting/common/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type TokenService struct {
	TokenRepository *TokenRepository
	UserService     *user.UserService
}

func NewTokenService(tokenRepository *TokenRepository, userService *user.UserService) *TokenService {
	return &TokenService{TokenRepository: tokenRepository, UserService: userService}
}

func (s *TokenService) GetToken(c *gin.Context) {
	s.TokenRepository.GetToken(c)
}

func (s *TokenService) CreateTokenForUser(user *user.User) (*Token, error) {
	latestUserDetails, err := s.UserService.GetUserByEmail(user.Email)
	if latestUserDetails == nil || err != nil {
		return nil, err
	}

	expirationTime, tokenString, err := utils.GenerateToken(latestUserDetails)
	if err != nil {
		return nil, err
	}

	token := &Token{
		UserId:    latestUserDetails.ID.ID,
		Token:     tokenString,
		TokenHash: utils.GenerateTokenHash(tokenString),
		ExpiresAt: expirationTime,
	}

	if err := s.CreateToken(token); err != nil {
		return nil, err
	}

	return token, nil
}

func (s *TokenService) CreateToken(token *Token) error {
	return s.TokenRepository.CreateToken(token)
}

func (s *TokenService) FindByTokenHash(tokenHash string) (*Token, error) {
	return s.TokenRepository.FindByTokenHash(tokenHash)
}

func (s *TokenService) DeleteToken(tokenHash string) (bool, error) {
	return s.TokenRepository.DeleteToken(tokenHash)
}

func (s *TokenService) DeleteTokenByTokenHash(tokenHash string) (bool, error) {
	return s.TokenRepository.DeleteTokenByTokenHash(tokenHash)
}

func (s *TokenService) UpdateToken(token *Token) error {
	return s.TokenRepository.UpdateToken(token)
}

func (s *TokenService) isValidToken(tokenString string) bool {
	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		return false
	}

	tokenHash := utils.GenerateTokenHash(tokenString)
	_, err = s.FindByTokenHash(tokenHash)
	if err != nil {
		return false
	}

	return claims.ExpiresAt > time.Now().Unix()
}

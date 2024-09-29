package token

import (
	"net/http"
	"next-gen-job-hunting/api/user"
	"next-gen-job-hunting/common/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TokenRepository struct {
	DB             *gorm.DB
	UserRepository *user.UserRepository
}

func NewTokenRepository(db *gorm.DB, userRepository *user.UserRepository) *TokenRepository {
	return &TokenRepository{DB: db, UserRepository: userRepository}
}

func (r *TokenRepository) GetToken(c *gin.Context) {
	tokenHash := c.Param("token_hash")
	token, err := r.FindByTokenHash(tokenHash)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Token not found"})
		return
	}
	c.JSON(http.StatusOK, token)
}

func (r *TokenRepository) CreateTokenForUser(user *user.User, c *gin.Context) (*Token, error) {
	latestUserDetails, err := r.UserRepository.GetUserByEmail(user.Email, c)
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
	if err := r.CreateToken(token); err != nil {
		return nil, err
	}

	return token, nil
}

func (r *TokenRepository) CreateToken(token *Token) error {
	return r.DB.Create(token).Error
}

func (r *TokenRepository) FindByTokenHash(tokenHash string) (*Token, error) {
	var token Token
	if err := r.DB.Where("token_hash = ?", tokenHash).First(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *TokenRepository) DeleteToken(tokenHash string) (bool, error) {
	if err := r.DB.Where("token_hash = ?", tokenHash).Delete(&Token{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *TokenRepository) DeleteTokenByTokenHash(tokenHash string) (bool, error) {
	if err := r.DB.Where("token_hash = ?", tokenHash).Delete(&Token{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *TokenRepository) UpdateToken(token *Token) error {
	return r.DB.Save(token).Error
}

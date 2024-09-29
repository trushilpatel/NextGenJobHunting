package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"next-gen-job-hunting/api/user"
	"next-gen-job-hunting/config/env"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(env.GetJWTSecret())

type Claims struct {
	UserId uint   `json:"userId"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func GenerateTokenHash(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func GenerateToken(user *user.User) (time.Time, string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserId: user.ID.ID,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	return expirationTime, tokenString, err
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

// ValidateToken validates a JWT token and returns the user ID if valid
func ValidateToken(tokenString string) (*Claims, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

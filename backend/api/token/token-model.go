package token

import (
	"next-gen-job-hunting/api/user"
	"next-gen-job-hunting/common/db"
	"time"
)

// Token represents a token model in the database.
//
// swagger:model Token
//
// Fields:
//   - UserId: The ID of the user associated with the token. This field is required.
//     example: 123
//   - Token: The token string. This field is required and indexed.
//     example: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
//   - TokenHash: The hashed version of the token. This field is required.
//   - ExpiresAt: The expiration time of the token. This field is required.
//     example: "2023-12-31T23:59:59Z"
//   - User: The user associated with the token. This field is not serialized in JSON responses.
type Token struct {
	db.ID
	UserId    uint      `gorm:"not null" json:"userId"`
	Token     string    `gorm:"type:varchar(500);not null;index" json:"token"`
	TokenHash string    `gorm:"type:varchar(64);not null" json:"_"`
	ExpiresAt time.Time `gorm:"not null" json:"expiresAt"`
	User      user.User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	db.CreatedAt
}

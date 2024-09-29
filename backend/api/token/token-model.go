package token

import (
	"next-gen-job-hunting/api/user"
	"next-gen-job-hunting/common/db"
	"time"
)

type Token struct {
	db.ID
	UserId    uint      `gorm:"not null" json:"userId"`
	Token     string    `gorm:"type:varchar(500);not null;index" json:"token"`
	TokenHash string    `gorm:"type:varchar(64);not null" json:"_"`
	ExpiresAt time.Time `gorm:"not null" json:"expiresAt"`
	User      user.User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	db.CreatedAt
}

package db

import (
	"time"
)

type ID struct {
	ID uint `gorm:"primaryKey" json:"id"`
}

type IdCreatedUpdated struct {
	ID
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

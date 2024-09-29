package db

import (
	"time"
)

type ID struct {
	ID uint `gorm:"primaryKey" json:"id"`
}

type CreatedAt struct {
	CreatedAt time.Time `json:"created_at"`
}

type UpdatedAt struct {
	UpdatedAt time.Time `json:"updated_at"`
}

type IdCreatedUpdated struct {
	ID
	CreatedAt
	UpdatedAt
}

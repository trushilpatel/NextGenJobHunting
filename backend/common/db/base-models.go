package db

import (
	"time"
)

// ID represents a model with a primary key ID field.
// The ID field is of type uint and is annotated to be the primary key in the database
// and to be serialized as "id" in JSON format.
//
// swagger:model ID
type ID struct {
	ID uint `gorm:"primaryKey" json:"id"`
}

// CreatedAt represents a model that includes a timestamp indicating when the record was created.
//
// swagger:model CreatedAt
//
// Fields:
// - CreatedAt: The timestamp when the record was created, formatted as a JSON date-time string.
type CreatedAt struct {
	CreatedAt time.Time `json:"created_at"`
}

// UpdatedAt represents a model that includes a timestamp indicating when the record was last updated.
// This field is serialized to JSON as "updated_at".
//
// swagger:model UpdatedAt
type UpdatedAt struct {
	UpdatedAt time.Time `json:"updated_at"`
}

// IdCreatedUpdated is a base model that includes fields for ID, creation timestamp, and update timestamp.
// swagger:model IdCreatedUpdated
type IdCreatedUpdated struct {
	ID
	CreatedAt
	UpdatedAt
}

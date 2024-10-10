package joburl

import (
	"next-gen-job-hunting/api/user"
	"next-gen-job-hunting/common/db"

	"gorm.io/gorm"
)

// Job Priorities
const (
	PriorityHigh    = "high"
	PriorityMedium  = "medium"
	PriorityLow     = "low"
	StatusNew       = "new"
	StatusCompleted = "completed"
)

// JobUrl represents a job URL entity with associated metadata.
//
// swagger:model JobUrl
//
// Fields:
//   - URL: The URL of the job posting. It has a maximum length of 2048 characters and cannot be null.
//     example: "https://example.com/job-posting"
//   - Priority: The priority of the job URL. It has a maximum length of 10 characters, defaults to 'low', and cannot be null.
//     example: "high"
//   - Status: The status of the job URL. It has a maximum length of 10 characters, defaults to 'new', and cannot be null.
//     example: "pending"
//   - UserID: The ID of the user associated with this job URL. It cannot be null.
//     example: 123
//   - User: The user entity associated with this job URL. This field is used to define the foreign key relationship and is not included in the JSON representation.
type JobUrl struct {
	db.IdCreatedUpdated
	URL      string    `gorm:"size:2048;not null" json:"url"`                  // URL field with a max length of 2048 characters
	Priority string    `gorm:"size:10;not null;default:'low'" json:"priority"` // Default priority set to 'low'
	Status   string    `gorm:"size:10;not null;default:'new'" json:"status"`   // Status field, e.g., "pending", "completed"
	UserID   uint      `gorm:"not null" json:"userId"`
	User     user.User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"` // Define foreign key relationship
}

// BeforeSave hook to validate and set default values
func (job *JobUrl) BeforeSave(tx *gorm.DB) (err error) {
	// Ensure the priority value is valid
	if job.Priority != PriorityHigh && job.Priority != PriorityMedium && job.Priority != PriorityLow {
		job.Priority = PriorityLow
	}

	// Ensure the status value is valid
	if job.Status != StatusNew && job.Status != StatusCompleted {
		job.Status = StatusNew
	}
	return nil
}

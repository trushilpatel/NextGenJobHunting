package jobs

import (
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

type JobUrl struct {
	db.IdCreatedUpdated
	URL      string `gorm:"size:2048;not null" json:"url"`                  // URL field with a max length of 2048 characters
	Priority string `gorm:"size:10;not null;default:'low'" json:"priority"` // Default priority set to 'l'
	Status   string `gorm:"size:10;not null;default:'new'" json:"status"`   // Status field, e.g., "pending", "completed"
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

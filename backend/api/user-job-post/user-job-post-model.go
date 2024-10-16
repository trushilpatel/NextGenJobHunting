package user_job_post

import (
	"next-gen-job-hunting/api/user"
	"next-gen-job-hunting/common/db"
	"time"
)

type JobApplicationStatus string

const (
	Saved     JobApplicationStatus = "Saved"
	Applied   JobApplicationStatus = "Applied"
	Interview JobApplicationStatus = "Interview"
	Offered   JobApplicationStatus = "Offered"
	Rejected  JobApplicationStatus = "Rejected"
	Withdrawn JobApplicationStatus = "Withdrawn"
)

var UserJobPostScripts = []string{
	`DO $$ 
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'jobapplicationstatus') THEN
			CREATE TYPE jobapplicationstatus AS ENUM ('Saved', 'Applied', 'Interview', 'Offered', 'Rejected', 'Withdrawn');
		END IF;
	END$$;`,
}

// UserJobPost represents the association between a user and a job post.
//
// swagger:model UserJobPost
//
// Fields:
//   - JobPostId: Foreign key to the job_post table.
//     required: true
//     example: 1
//   - UserId: Foreign key to the user table.
//     required: true
//     example: 1
//   - IsEligible: Indicates whether the user is eligible for the job.
//     required: false
//     example: false
//   - JobApplicationStatus: Status of the job application.
//     required: true
//     example: Saved
//   - AppliedAt: Timestamp when the user applied for the job.
//     required: false
//     example: 2023-01-01T00:00:00Z
//   - ResumeId: ID referencing the user's resume.
//     required: false
//     example: 1
//   - CoverLetterId: ID referencing the user's cover letter.
//     required: false
//     example: 1
//   - IsStatusHidden: Flag to indicate if the user wants to hide the job application status.
//     required: false
//     example: false
//   - ResumeScore: Score given by an ATS for the resume.
//     required: false
//     example: 85
//   - CreatedAt: Timestamp when the record was created.
//     required: true
//     example: 2023-01-01T00:00:00Z
//   - UpdatedAt: Timestamp when the record was last updated.
//     required: true
//     example: 2023-01-01T00:00:00Z
//   - User: The user associated with this job post.
//     required: true
type UserJobPost struct {
	JobPostId            uint                 `gorm:"not null" json:"jobPostId"`                                           // Foreign key to job_post table
	UserId               uint                 `gorm:"not null" json:"userId"`                                              // Foreign key to user table
	IsEligible           bool                 `gorm:"default:false" json:"isEligible"`                                     // Whether the user is eligible for the job
	JobApplicationStatus JobApplicationStatus `gorm:"type:varchar(20);not null;default:Saved" json:"jobApplicationStatus"` // Status of the job application
	AppliedAt            time.Time            `json:"appliedAt"`                                                           // Timestamp when the user applied for the job
	ResumeId             uint                 `json:"resumeId"`                                                            // ID referencing the user's resume
	CoverLetterId        uint                 `json:"coverLetterId"`                                                       // ID referencing the user's cover letter
	IsStatusHidden       bool                 `gorm:"default:false" json:"isStatusHidden"`                                 // Flag to indicate if the user wants to hide the job application status
	ResumeScore          int                  `json:"resumeScore"`                                                         // Score given by an ATS for the resume
	db.CreatedAt
	db.UpdatedAt
	User user.User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	// there is a circular dependency between user_job_post and job_post how ever there is no need to import job_post here
	// we just need a relation to job_post
	//JobPost jobpost.JobPost `gorm:"foreignKey:JobPostId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

}

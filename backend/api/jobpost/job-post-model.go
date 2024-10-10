package jobpost

import (
	userjobpost "next-gen-job-hunting/api/user-job-post"
	"next-gen-job-hunting/common/db"
	"time"
)

type JobType string

const (
	OnSite JobType = "On-site"
	Remote JobType = "Remote"
	Hybrid JobType = "Hybrid"
)

type EmploymentType string

const (
	FullTime   EmploymentType = "Full-time"
	PartTime   EmploymentType = "Part-time"
	Contract   EmploymentType = "Contract"
	Internship EmploymentType = "Internship"
)

var SqlScripts = []string{
	`DO $$ BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'job_type') THEN
            CREATE TYPE job_type AS ENUM ('On-site', 'Remote', 'Hybrid');
        END IF;
    END $$;`,
	`DO $$ BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'employment_type') THEN
            CREATE TYPE employment_type AS ENUM ('Full-time', 'Part-time', 'Contract', 'Internship');
        END IF;
    END $$;`,
}

// JobPost represents a job posting in the system.
//
// swagger:model JobPost
type JobPost struct {
	db.ID
	// Unique identifier for the job post
	// required: true
	JobID string `gorm:"unique;not null;size:50" json:"jobId"` // Unique Job ID, up to 50 characters
	// Title of the job
	// required: true
	JobTitle string `gorm:"not null;size:50" json:"jobTitle"` // Job title, up to 50 characters
	// Type of employment (e.g., full-time, part-time)
	// required: true
	EmploymentType EmploymentType `gorm:"type:varchar(10);not null" json:"employmentType"` // Enum for employment type
	// Salary range for the job
	SalaryRange string `gorm:"size:50" json:"salaryRange"` // Salary range, up to 50 characters
	// Type of job (e.g., contract, freelance)
	// required: true
	JobType JobType `gorm:"type:varchar(10);not null" json:"jobType"` // Enum for job type
	// Duration of the contract, if applicable
	ContractLength string `gorm:"size:50" json:"contractLength"` // Contract duration, if applicable, up to 50 characters
	// Industry related to the job
	Industry string `gorm:"size:50" json:"industry"` // Industry related to the job, up to 50 characters
	// Skills required for the job
	RequiredSkills string `gorm:"size:100" json:"requiredSkills"` // Skills required for the job, up to 100 characters
	// Minimum education level required for the job
	EducationLevel string `gorm:"size:50" json:"educationLevel"` // Minimum education level required, up to 50 characters
	// Date when the job was posted
	// required: true
	JobPostedDate time.Time `gorm:"not null" json:"jobPostedDate"` // Date the job was posted
	// Detailed description of the job
	// required: true
	JobDescription string `gorm:"size:1000;not null" json:"jobDescription"` // Detailed description of the job, up to 1000 characters
	// URL where applicants can apply for the job
	// required: true
	ApplicationLink string `gorm:"size:200;not null" json:"applicationLink"` // URL for applying to the job, up to 200 characters
	// Indicates if the person must be a U.S. citizen or permanent resident
	IsRequireUSAPerson bool `json:"isRequireUSAPerson"` // Indicates if the person must be a U.S. citizen or permanent resident, default: false
	// Indicates if security clearance is required
	IsSecurityClearanceRequired bool `json:"isSecurityClearanceRequired"` // Whether security clearance is required, default: false
	// Visa sponsorship information, if any
	VisaSponsorshipDetails string `gorm:"size:50" json:"visaSponsorshipDetails"` // Visa sponsorship information, if any, up to 50 characters
	// Minimum experience required for the job
	MinimumExperienceRequired int `json:"minimumExperienceRequired"` // Minimum experience required for the job
	// Number of applicants for the job
	Applicants int `gorm:"default:0" json:"applicants"` // Number of applicants
	// The company or person responsible for hiring
	Hirer string `gorm:"size:30" json:"hirer"` // The company/person responsible for hiring, up to 30 characters
	// Profile link of the hirer (e.g., LinkedIn, company page)
	HirerProfileLink string `gorm:"size:100" json:"hirerProfileLink"` // Profile link of the hirer, up to 100 characters
	// Location of the job
	Location string `gorm:"size:30" json:"location"` // Job location, up to 30 characters
	// Keywords for the application tracking system
	AtsKeywords string `gorm:"size:200" json:"atsKeywords"` // ATS keywords, up to 200 characters
	// Automatically sets the creation timestamp
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"` // Automatically sets the creation timestamp
}

type JobPostUserJobPostDto struct {
	JobPost
	userjobpost.UserJobPost
}

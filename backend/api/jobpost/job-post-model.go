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

type JobPost struct {
	db.ID
	JobID                       string         `gorm:"unique;not null;size:50" json:"jobId"`            // Unique Job ID, up to 30 characters
	JobTitle                    string         `gorm:"not null;size:50" json:"jobTitle"`                // Job title, up to 30 characters
	EmploymentType              EmploymentType `gorm:"type:varchar(10);not null" json:"employmentType"` // Enum for employment type
	SalaryRange                 string         `gorm:"size:50" json:"salaryRange"`                      // Salary range, up to 30 characters
	JobType                     JobType        `gorm:"type:varchar(10);not null" json:"jobType"`        // Enum for job type
	ContractLength              string         `gorm:"size:50" json:"contractLength"`                   // Contract duration, if applicable, up to 30 characters
	Industry                    string         `gorm:"size:50" json:"industry"`                         // Industry related to the job, up to 30 characters
	RequiredSkills              string         `gorm:"size:100" json:"requiredSkills"`                  // Skills required for the job (can be a JSON list or a comma-separated string), up to 100 characters
	EducationLevel              string         `gorm:"size:50" json:"educationLevel"`                   // Minimum education level required, up to 30 characters
	JobPostedDate               time.Time      `gorm:"not null" json:"jobPostedDate"`                   // Date the job was posted
	JobDescription              string         `gorm:"size:1000;not null" json:"jobDescription"`        // Detailed description of the job, up to 1000 characters
	ApplicationLink             string         `gorm:"size:200;not null" json:"applicationLink"`        // URL for applying to the job, up to 200 characters
	IsRequireUSAPerson          bool           `json:"isRequireUSAPerson"`                              // Indicates if the person must be a U.S. citizen or permanent resident, default: false
	IsSecurityClearanceRequired bool           `json:"isSecurityClearanceRequired"`                     // Whether security clearance is required, default: false
	VisaSponsorshipDetails      string         `gorm:"size:50" json:"visaSponsorshipDetails"`           // Visa sponsorship information, if any, up to 50 characters
	MinimumExperienceRequired   int            `json:"minimumExperienceRequired"`                       // Minimum experience required for the job
	Applicants                  int            `gorm:"default:0" json:"applicants"`                     // Number of applicants
	Hirer                       string         `gorm:"size:30" json:"hirer"`                            // The company/person responsible for hiring, up to 20 characters
	HirerProfileLink            string         `gorm:"size:100" json:"hirerProfileLink"`                // Profile link of the hirer (e.g., LinkedIn, company page), up to 100 characters
	Location                    string         `gorm:"size:30" json:"location"`                         // Job location, up to 30 characters
	AtsKeywords                 string         `gorm:"size:200" json:"atsKeywords"`                     // application tracking system keywords (comma separated) to get top 100% resume ATS score
	CreatedAt                   time.Time      `gorm:"autoCreateTime" json:"createdAt"`                 // Automatically sets the creation timestamp
}

type JobPostUserJobPostDto struct {
	JobPost
	userjobpost.UserJobPost
}

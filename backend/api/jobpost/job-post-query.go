package jobpost

import (
	"next-gen-job-hunting/api/common"
)

// JobPostQuery holds the filter and pagination params for the job post search
type JobPostQuery struct {
	common.Pagination           // Embed the common pagination struct
	JobPostId            uint   `form:"jobPostId"`            // JobPost ID for filtering
	UserId               uint   `form:"userId"`               // User ID for filtering
	IsEligible           bool   `form:"isEligible"`           // Eligibility flag
	IsRequireUSAPerson   bool   `form:"isRequireUSAPerson"`   // USA person requirement flag
	JobApplicationStatus string `form:"jobApplicationStatus"` // Job application status filter
	Hirer                string `form:"hirer"`                // Hirer name for filtering
	Location             string `form:"location"`             // Job location filter
}

// Validate validates the JobPostQuery and sets default values if necessary
func (q *JobPostQuery) Validate() {
	q.Pagination.Validate()
	if q.Pagination.SortBy == "" || (q.Pagination.SortBy != "CreatedAt" && q.Pagination.SortBy != "JobPostedDate" && q.Pagination.SortBy != "Applicants") {
		q.Pagination.SortBy = "id"
	}
}

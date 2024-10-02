package jobpost

import (
	"next-gen-job-hunting/api/common"
	userjobpost "next-gen-job-hunting/api/user-job-post"
)

// JobPostQuery holds the filter and pagination params for the job post search
type JobPostQuery struct {
	common.Pagination                                     // Embed the common pagination struct
	JobPostId            uint                             `form:"jobPostId"`            // JobPost ID for filtering
	UserId               uint                             `form:"userId"`               // User ID for filtering
	IsEligible           bool                             `form:"isEligible"`           // Eligibility flag
	IsRequireUSAPerson   bool                             `form:"isRequireUSAPerson"`   // USA person requirement flag
	JobApplicationStatus userjobpost.JobApplicationStatus `form:"jobApplicationStatus"` // Job application status filter
	Hirer                string                           `form:"hirer"`                // Hirer name for filtering
	Location             string                           `form:"location"`             // Job location filter
}

// Validate validates the JobPostQuery and sets default values if necessary
func (q *JobPostQuery) Validate() {
	q.Pagination.Validate()
	if q.Pagination.SortBy == "" || (q.Pagination.SortBy != "CreatedAt" && q.Pagination.SortBy != "JobPostedDate" && q.Pagination.SortBy != "Applicants") {
		q.Pagination.SortBy = "id"
	}
	if q.JobApplicationStatus != "" &&
		q.JobApplicationStatus != userjobpost.Saved &&
		q.JobApplicationStatus != userjobpost.Applied &&
		q.JobApplicationStatus != userjobpost.Interview &&
		q.JobApplicationStatus != userjobpost.Offered &&
		q.JobApplicationStatus != userjobpost.Rejected &&
		q.JobApplicationStatus != userjobpost.Withdrawn {
		q.JobApplicationStatus = ""
	}
}

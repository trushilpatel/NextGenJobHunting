package jobpost

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JobPostService struct {
	Repo *JobPostRepository
}

func NewJobPostService(repo *JobPostRepository) *JobPostService {
	return &JobPostService{Repo: repo}
}

func (s *JobPostService) Create(jobPost *JobPost, c *gin.Context) error {
	return s.Repo.Create(jobPost, c)
}

func (s *JobPostService) Search(query JobPostQuery, c *gin.Context) ([]JobPost, error) {
	db := s.Repo.DB

	fmt.Printf("JobPostId: %d\n", query.JobPostId)
	fmt.Printf("UserId: %d\n", query.UserId)
	fmt.Printf("IsEligible: %t\n", query.IsEligible)
	fmt.Printf("IsRequireUSAPerson: %t\n", query.IsRequireUSAPerson)
	fmt.Printf("JobApplicationStatus: %s\n", query.JobApplicationStatus)
	fmt.Printf("Hirer: %s\n", query.Hirer)
	fmt.Printf("Location: %s\n", query.Location)

	// Apply pagination before building the query

	db = query.Pagination.ApplyToDB(db)

	// Apply filters
	if query.JobPostId != 0 {
		jobPost, err := s.FindByID(query.JobPostId, c)
		if err != nil {
			return nil, err
		}
		return []JobPost{*jobPost}, nil
	}
	if query.UserId != 0 {
		// You will have to implement this part
		// db = db.Where("user_id = ?", query.UserId)
	}
	if query.IsEligible {
		// You will have to implement this part as well
		// db = db.Where("is_eligible = ?", query.IsEligible)
	}
	if query.IsRequireUSAPerson {
		db = db.Where("is_require_usa_person = ?", query.IsRequireUSAPerson)
	}
	// if query.JobApplicationStatus != "" {
	// 	db = db.Where("job_application_status = ?", query.JobApplicationStatus)
	// }
	if query.Hirer != "" {
		db = db.Where("hirer LIKE ?", "%"+query.Hirer+"%")
	}
	if query.Location != "" {
		db = db.Where("location LIKE ?", "%"+query.Location+"%")
	}

	// Generate the SQL query string for debugging purposes
	queryStr := db.Model(&JobPost{}).Find(&[]JobPost{}).ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx
	})

	// Print the final SQL query
	fmt.Printf("Final DB Query: %s\n", queryStr)

	// Now execute the query only once
	var jobPosts []JobPost
	if err := db.Find(&jobPosts).Error; err != nil {
		return nil, err
	}

	return jobPosts, nil
}

func (s *JobPostService) FindAll(c *gin.Context) ([]JobPost, error) {
	return s.Repo.FindAll(c)
}

func (s *JobPostService) FindByID(id uint, c *gin.Context) (*JobPost, error) {
	return s.Repo.FindByID(id, c)
}

func (s *JobPostService) Update(jobPost *JobPost, c *gin.Context) error {
	return s.Repo.Update(jobPost, c)
}

func (s *JobPostService) Delete(id uint, c *gin.Context) error {
	return s.Repo.Delete(id, c)
}

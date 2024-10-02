package jobpost

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JobPostRepository struct {
	DB *gorm.DB
}

func NewJobPostRepository(db *gorm.DB) *JobPostRepository {
	return &JobPostRepository{DB: db}
}

func (r *JobPostRepository) Create(jobPost *JobPost, c *gin.Context) error {
	return r.DB.Create(jobPost).Error
}
func (s *JobPostRepository) Search(query JobPostQuery, c *gin.Context) ([]JobPostUserJobPostDto, error) {
	db := s.DB

	fmt.Printf("JobPostId: %d\n", query.JobPostId)
	fmt.Printf("UserId: %d\n", query.UserId)
	fmt.Printf("IsEligible: %t\n", query.IsEligible)
	fmt.Printf("IsRequireUSAPerson: %t\n", query.IsRequireUSAPerson)
	fmt.Printf("JobApplicationStatus: %s\n", query.JobApplicationStatus)
	fmt.Printf("Hirer: %s\n", query.Hirer)
	fmt.Printf("Location: %s\n", query.Location)

	// Apply pagination before building the query
	db = query.Pagination.ApplyToDB(db)

	// Start with the base query, join UserJobPost on JobPostId
	db = db.Table("job_post").Select("job_post.*, ujp.*").
		Joins("LEFT JOIN user_job_post ujp ON ujp.job_post_id = job_post.id")

	// Apply filters
	if query.JobPostId != 0 {
		db = db.Where("job_post.id = ?", query.JobPostId)
	}
	if query.UserId != 0 {
		db = db.Where("ujp.user_id = ?", query.UserId)
	}
	if query.IsEligible {
		db = db.Where("ujp.is_eligible = ?", query.IsEligible)
	}
	if query.IsRequireUSAPerson {
		db = db.Where("job_post.is_require_usa_person = ?", query.IsRequireUSAPerson)
	}
	if query.Hirer != "" {
		db = db.Where("job_post.hirer LIKE ?", "%"+query.Hirer+"%")
	}
	if query.Location != "" {
		db = db.Where("job_post.location LIKE ?", "%"+query.Location+"%")
	}
	if query.JobApplicationStatus != "" {
		db = db.Where("ujp.job_application_status = ?", query.JobApplicationStatus)
	}

	// Execute the query
	var jobPosts []JobPostUserJobPostDto
	if err := db.Scan(&jobPosts).Error; err != nil {
		return nil, err
	}

	return jobPosts, nil
}

func (r *JobPostRepository) FindAll(c *gin.Context) ([]JobPost, error) {
	var jobPosts []JobPost
	err := r.DB.Find(&jobPosts).Error
	return jobPosts, err
}

func (r *JobPostRepository) FindByID(id uint, c *gin.Context) (*JobPost, error) {
	var jobPost JobPost
	err := r.DB.First(&jobPost, id).Error
	return &jobPost, err
}

func (r *JobPostRepository) Update(jobPost *JobPost, c *gin.Context) error {
	return r.DB.Save(jobPost).Error
}

func (r *JobPostRepository) Delete(id uint, c *gin.Context) error {
	return r.DB.Delete(&JobPost{}, id).Error
}

package joburl

import (
	"gorm.io/gorm"
)

type JobUrlRepository struct {
	DB *gorm.DB
}

func NewJobUrlRepository(db *gorm.DB) *JobUrlRepository {
	return &JobUrlRepository{DB: db}
}

func (r *JobUrlRepository) CreateJob(jobUrl *JobUrl) error {
	if err := r.DB.Create(jobUrl).Error; err != nil {
		return err
	}
	return nil
}

func (r *JobUrlRepository) GetAllJobUrl() ([]*JobUrl, error) {
	var jobUrls []*JobUrl
	if err := r.DB.Find(&jobUrls).Error; err != nil {
		return nil, err
	}
	return jobUrls, nil
}

func (r *JobUrlRepository) GetJobUrlById(id uint) (*JobUrl, error) {
	var jobUrl JobUrl
	if err := r.DB.First(&jobUrl, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &jobUrl, nil
}

func (r *JobUrlRepository) UpdateJobUrl(jobUrl *JobUrl) error {
	if err := r.DB.Save(jobUrl).Error; err != nil {
		return err
	}
	return nil
}

func (r *JobUrlRepository) DeleteJobUrl(id uint) error {
	if err := r.DB.Delete(&JobUrl{}, id).Error; err != nil {
		return err
	}
	return nil
}

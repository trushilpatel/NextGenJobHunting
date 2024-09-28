package joburl

type JobUrlService struct {
	repo *JobUrlRepository
}

func NewJobUrlService(repo *JobUrlRepository) *JobUrlService {
	return &JobUrlService{repo: repo}
}

func (s *JobUrlService) CreateJobUrl(jobUrl *JobUrl) error {
	return s.repo.CreateJob(jobUrl)
}

func (s *JobUrlService) GetAllJobUrl() ([]*JobUrl, error) {
	return s.repo.GetAllJobUrl()
}

func (s *JobUrlService) GetJobUrlById(id uint) (*JobUrl, error) {
	return s.repo.GetJobUrlById(id)
}

func (s *JobUrlService) UpdateJobUrl(jobUrl *JobUrl) error {
	return s.repo.UpdateJobUrl(jobUrl)
}

func (s *JobUrlService) DeleteJobUrl(id uint) error {
	return s.repo.DeleteJobUrl(id)
}

package application

import "github.com/bug-free-happiness/easy-issues/domain"

type IssueService struct {
	IssueRepository domain.IssueRepository
}

// Create implements domain.IssueService.
func (s *IssueService) Create(issue *domain.Issue) error {
	return s.IssueRepository.Create(issue)
}

// Delete implements domain.IssueService.
func (s *IssueService) Delete(id int64) error {
	return s.IssueRepository.Delete(id)
}

// Issue implements domain.IssueService.
func (s *IssueService) Issue(id int64) (*domain.Issue, error) {
	return s.IssueRepository.GetById(id)
}

// Issues implements domain.IssueService.
func (s *IssueService) Issues() ([]*domain.Issue, error) {
	return s.IssueRepository.All()
}

func NewIssueService(issueRepository domain.IssueRepository) domain.IssueService {
	return &IssueService{issueRepository}
}

package application

import "github.com/bug-free-happiness/easy-issues/domain"

type ProjectService struct {
	ProjectRepository domain.ProjectRepository
}

// Create implements domain.ProjectService.
func (s *ProjectService) Create(p *domain.Project) error {
	return s.ProjectRepository.Create(p)
}

// Delete implements domain.ProjectService.
func (s *ProjectService) Delete(id int64) error {
	return s.ProjectRepository.Delete(id)
}

// Project implements domain.ProjectService.
func (s *ProjectService) Project(id int64) (*domain.Project, error) {
	return s.ProjectRepository.GetById(id)
}

// Projects implements domain.ProjectService.
func (s *ProjectService) Projects() ([]*domain.Project, error) {
	return s.ProjectRepository.All()
}

func NewProjectService(projectRepository domain.ProjectRepository) domain.ProjectService {
	return &ProjectService{projectRepository}
}

package memory

import (
	"errors"

	"github.com/bug-free-happiness/easy-issues/domain"
	"github.com/patrickmn/go-cache"
)

const (
	ProjectAllKey = "Projects:all"
	ProjectLastId = "Project:lastId"
)

type ProjectRepository struct {
	db *cache.Cache
}

// All implements domain.ProjectRepository.
func (r *ProjectRepository) All() ([]*domain.Project, error) {
	result, ok := r.db.Get(ProjectAllKey)
	if ok {
		return result.([]*domain.Project), nil
	} else {
		return nil, errors.New("empty list")
	}
}

// Create implements domain.ProjectRepository.
func (r *ProjectRepository) Create(p *domain.Project) error {
	id, _ := r.db.IncrementInt64(ProjectLastId, int64(1))
	p.Id = id

	result, ok := r.db.Get(ProjectAllKey)
	if ok {
		result = append(result.([]*domain.Project), p)
		r.db.Set(ProjectAllKey, result, cache.NoExpiration)
	}

	return nil
}

// Delete implements domain.ProjectRepository.
func (r *ProjectRepository) Delete(id int64) error {
	result, ok := r.db.Get(ProjectAllKey)
	if ok {
		items := result.([]*domain.Project)
		for i, project := range items {
			if project.Id == id {
				items = append(items[:i], items[i+1:]...)
				r.db.Set(ProjectAllKey, items, cache.NoExpiration)
				return nil
			}
		}
		return errors.New("not found")
	}
	return errors.New("not found")
}

// GetById implements domain.ProjectRepository.
func (r *ProjectRepository) GetById(id int64) (*domain.Project, error) {
	result, ok := r.db.Get(ProjectAllKey)
	if ok {
		items := result.([]*domain.Project)
		for _, project := range items {
			if project.Id == id {
				return project, nil
			}
		}
		return nil, errors.New("not found")
	}
	return nil, errors.New("not found")
}

func NewProjectRepository() domain.ProjectRepository {
	db := cache.New(cache.NoExpiration, cache.NoExpiration)
	db.SetDefault(ProjectLastId, int64(0))
	db.SetDefault(ProjectAllKey, []*domain.Project{})

	return &ProjectRepository{db}
}

package memory

import (
	"errors"

	"github.com/bug-free-happiness/easy-issues/domain"
	"github.com/patrickmn/go-cache"
)

const (
	IssuesAllKey = "Issues:all"
	IssueLastId  = "Issue:lastId"
)

type IssueRepository struct {
	db *cache.Cache
}

// All implements domain.IssueRepository.
func (r *IssueRepository) All() ([]*domain.Issue, error) {
	result, ok := r.db.Get(IssuesAllKey)
	if ok {
		return result.([]*domain.Issue), nil
	} else {
		return nil, errors.New("empty list")
	}
}

// Create implements domain.IssueRepository.
func (r *IssueRepository) Create(issue *domain.Issue) error {
	id, _ := r.db.IncrementInt64(IssueLastId, int64(1))
	issue.Id = id

	result, ok := r.db.Get(IssuesAllKey)
	if ok {
		result = append(result.([]*domain.Issue), issue)
		r.db.Set(IssuesAllKey, result, cache.NoExpiration)
	}

	return nil
}

// Delete implements domain.IssueRepository.
func (r *IssueRepository) Delete(id int64) error {
	result, ok := r.db.Get(IssuesAllKey)
	if ok {
		items := result.([]*domain.Issue)
		for i, issue := range items {
			if issue.Id == id {
				items = append(items[:i], items[i+1:]...)
				r.db.Set(IssuesAllKey, items, cache.NoExpiration)
				return nil
			}
		}
		return errors.New("not found")
	}
	return errors.New("not found")
}

// GetById implements domain.IssueRepository.
func (r *IssueRepository) GetById(id int64) (*domain.Issue, error) {
	result, ok := r.db.Get(IssuesAllKey)
	if ok {
		items := result.([]*domain.Issue)
		for _, issue := range items {
			if issue.Id == id {
				return issue, nil
			}
		}
		return nil, errors.New("not found")
	}
	return nil, errors.New("not found")
}

func NewIssueRepository() domain.IssueRepository {
	db := cache.New(cache.NoExpiration, cache.NoExpiration)
	db.SetDefault(IssueLastId, int64(0))
	db.SetDefault(IssuesAllKey, []*domain.Issue{})

	return &IssueRepository{db}
}

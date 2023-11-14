package memory

import (
	"errors"

	"github.com/bug-free-happiness/easy-issues/domain"
	"github.com/patrickmn/go-cache"
)

const (
	UsersAllKey = "Users:all"
	UserLastId  = "User:lastId"
)

type UserRepository struct {
	db *cache.Cache
}

// All implements domain.UserRepository.
func (r *UserRepository) All() ([]*domain.User, error) {
	result, ok := r.db.Get(UsersAllKey)
	if ok {
		return result.([]*domain.User), nil
	}
	return nil, errors.New("Empty")
}

// Create implements domain.UserRepository.
func (r *UserRepository) Create(u *domain.User) error {
	id, _ := r.db.IncrementInt64(UserLastId, int64(1))
	u.Id = id

	result, ok := r.db.Get(UsersAllKey)
	if ok {
		result = append(result.([]*domain.User), u)
		r.db.Set(UsersAllKey, result, cache.NoExpiration)
	}

	return nil
}

// Delete implements domain.UserRepository.
func (r *UserRepository) Delete(id int64) error {
	result, ok := r.db.Get(UsersAllKey)
	if ok {
		items := result.([]*domain.User)
		for i, user := range items {
			if user.Id == id {
				items = append(items[:i], items[i+1:]...)
				r.db.Set(UsersAllKey, items, cache.NoExpiration)
				return nil
			}
		}
		return errors.New("not found")
	}

	return errors.New("not found")
}

// GetById implements domain.UserRepository.
func (r *UserRepository) GetById(id int64) (*domain.User, error) {
	result, ok := r.db.Get(UsersAllKey)
	if ok {
		items := result.([]*domain.User)
		for _, user := range items {
			if user.Id == id {
				return user, nil
			}
		}
		return nil, errors.New("empty")
	}
	return nil, errors.New("empty")
}

func NewUserRepository() domain.UserRepository {
	db := cache.New(cache.NoExpiration, cache.NoExpiration)
	db.SetDefault(UserLastId, int64(0))
	db.SetDefault(UsersAllKey, []*domain.User{})

	return &UserRepository{db}
}

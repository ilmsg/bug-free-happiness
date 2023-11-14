package application

import "github.com/bug-free-happiness/easy-issues/domain"

type UserService struct {
	UserRepository domain.UserRepository
}

// Create implements domain.UserService.
func (s *UserService) Create(u *domain.User) error {
	return s.UserRepository.Create(u)
}

// Delete implements domain.UserService.
func (s *UserService) Delete(id int64) error {
	return s.UserRepository.Delete(id)
}

// User implements domain.UserService.
func (s *UserService) User(id int64) (*domain.User, error) {
	return s.UserRepository.GetById(id)
}

// Users implements domain.UserService.
func (s *UserService) Users() ([]*domain.User, error) {
	return s.UserRepository.All()
}

func NewUserService(userRepository domain.UserRepository) domain.UserService {
	return &UserService{userRepository}
}

package domain

import (
	"context"
)

type authService struct {
	storage Storage
	hasher  Hasher
}

// Authenticate implements AuthService.
func (s *authService) Authenticate(ctx context.Context, r AuthParams) (AuthTokenID, error) {
	// a, err := s.storage.FindByName(ctx, r.Name)
	// if err != nil {
	// 	log.Printf("error: find credential by name %v\n", err)
	// 	return "", fmt.Errorf("find credential by name failed")
	// }

	// if !s.hasher.Compare(r.Password, a.PasswordHash)
	panic("unimplemented")
}

// Expire implements AuthService.
func (s *authService) Expire(ctx context.Context, id AuthTokenID) error {
	panic("unimplemented")
}

// Validate implements AuthService.
func (s *authService) Validate(ctx context.Context, id AuthTokenID) error {
	panic("unimplemented")
}

func NewAuthService(s Storage, h Hasher) AuthService {
	return &authService{storage: s, hasher: h}
}

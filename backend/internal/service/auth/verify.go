package auth

import "github.com/drizzleent/wallet-tracker/backend/internal/model"

func (s *service) Verify(tokenStr string) (*model.User, error) {
	claims, err := s.jwtProvider.Verify(tokenStr)
	if err != nil {
		return nil, err
	}

	u, err := s.repo.Get(claims.Subject)
	if err != nil {
		return nil, err
	}

	return u, nil
}

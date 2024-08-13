package auth

import (
	"context"

	"github.com/drizzleent/wallet-tracker/backend/internal/model"
)

func (s *service) UserNonce(ctx context.Context, id string) (*model.User, error) {
	return s.repo.UserNonce(id)
}

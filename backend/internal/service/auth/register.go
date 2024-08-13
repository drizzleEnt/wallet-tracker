package auth

import (
	"context"

	"github.com/drizzleent/wallet-tracker/backend/internal/model"
)

func (s *service) Register(ctx context.Context, p *model.RegisterPayload) (*model.User, error) {
	return s.repo.Register(p)
}

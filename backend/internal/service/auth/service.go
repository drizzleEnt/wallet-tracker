package auth

import (
	"time"

	"github.com/drizzleent/wallet-tracker/backend/internal/utils"
	"github.com/drizzleent/wallet-tracker/backend/repository"
)

type service struct {
	repo        repository.AuthRepository
	jwtProvider *utils.JwtHmacProvider
}

func NewService(r repository.AuthRepository) *service {
	return &service{
		repo:        r,
		jwtProvider: utils.NewJwtHmacProvider("env", "mm login", 2*time.Minute),
	}
}

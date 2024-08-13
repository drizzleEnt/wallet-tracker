package auth

import "github.com/drizzleent/wallet-tracker/backend/repository"

type service struct {
	repo repository.AuthRepository
}

func NewService(r repository.AuthRepository) *service {
	return &service{
		repo: r,
	}
}

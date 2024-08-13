package repository

import "github.com/drizzleent/wallet-tracker/backend/internal/model"

type AuthRepository interface {
	Register(*model.RegisterPayload) (*model.User, error)
	UserNonce(id string) (*model.User, error)
}

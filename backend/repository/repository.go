package repository

import "github.com/drizzleent/wallet-tracker/backend/internal/model"

type AuthRepository interface {
	Register(*model.RegisterPayload) (*model.User, error)
	UserNonce(string) (*model.User, error)
	Update(*model.User) error
	Get(string) (*model.User, error)
	Nonce() (string, error)
}

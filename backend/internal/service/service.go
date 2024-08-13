package service

import (
	"context"

	"github.com/drizzleent/wallet-tracker/backend/internal/model"
)

type AuthService interface {
	Register(context.Context, *model.RegisterPayload) (*model.User, error)
	Signin()
	UserNonce(context.Context, string) (*model.User, error)
	Welcome()
}

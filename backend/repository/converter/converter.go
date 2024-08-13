package converter

import (
	"github.com/drizzleent/wallet-tracker/backend/internal/model"
	datamodel "github.com/drizzleent/wallet-tracker/backend/repository/data_model"
)

func FromRegisterPayloadToDataUser(p *model.RegisterPayload) *datamodel.User {
	return &datamodel.User{
		Address: p.Address,
	}
}

func FromDataUserToModelUser(u *datamodel.User) *model.User {
	return &model.User{
		Address: u.Address,
		Nonce:   u.Nonce,
	}
}

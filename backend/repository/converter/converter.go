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
		Sig:     u.Sig,
	}
}

func FromUserToDataUser(u *model.User) *datamodel.User {
	return &datamodel.User{
		Address: u.Address,
		Nonce:   u.Nonce,
		Sig:     u.Sig,
	}
}

func FromSigPayloadToData(p *model.SigningPayload) *datamodel.User {
	return &datamodel.User{
		Address: p.Address,
		Nonce:   p.Nonce,
		Sig:     p.Sig,
	}
}

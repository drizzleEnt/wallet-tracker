package model

type RegisterPayload struct {
	Address string `json:"address"`
}

type SigningPayload struct {
	Address string `json:"address"`
	Nonce   string `json:"address"`
	Sig     string `json:"address"`
}

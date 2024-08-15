package auth

import (
	"errors"
	"strings"

	"github.com/drizzleent/wallet-tracker/backend/internal/model"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func (s *service) Signin(sp *model.SigningPayload) (string, error) {
	u, err := s.repo.Get(sp.Address)
	if err != nil {
		return "", err
	}

	if sp.Nonce != u.Nonce {
		return "", errors.New("authentication error")
	}

	sig := hexutil.MustDecode(sp.Sig)

	sig[crypto.RecoveryIDOffset] -= 27
	msg := accounts.TextHash([]byte(sp.Nonce))

	recovered, err := crypto.SigToPub(msg, sig)

	if err != nil {
		return "", err
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	if u.Address != strings.ToLower(recoveredAddr.Hex()) {
		return "", errors.New("authentication error")
	}

	nonce, err := s.repo.Nonce()
	if err != nil {
		return "", err
	}

	u.Nonce = nonce
	s.repo.Update(u)

	signedToken, err := s.jwtProvider.CreateStandart(u.Address)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

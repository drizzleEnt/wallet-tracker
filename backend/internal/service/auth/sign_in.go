package auth

import (
	"errors"
	"strings"

	"github.com/drizzleent/wallet-tracker/backend/internal/model"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func (s *service) Signin(sp *model.SigningPayload) (*model.User, error) {
	u, err := s.repo.Get(sp.Address)
	if err != nil {
		return nil, err
	}

	if sp.Nonce != u.Nonce {
		return nil, errors.New("authentication error")
	}

	sig := hexutil.MustDecode(sp.Sig)

	sig[crypto.RecoveryIDOffset] -= 27
	msg := accounts.TextHash([]byte(sp.Nonce))

	recovered, err := crypto.SigToPub(msg, sig)

	if err != nil {
		return nil, err
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	if u.Address != strings.ToLower(recoveredAddr.Hex()) {
		return nil, errors.New("authentication error")
	}

	nonce, err := s.repo.Nonce()
	if err != nil {
		return nil, err
	}

	u.Nonce = nonce
	s.repo.Update(u)

	return u, nil
}

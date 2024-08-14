package converter

import (
	"errors"
	"regexp"

	"github.com/drizzleent/wallet-tracker/backend/internal/model"
)

var (
	hexRegex   *regexp.Regexp = regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)
	nonceRegex *regexp.Regexp = regexp.MustCompile(`0[0-9]+$`)
)

func ValidateAddress(adrr string) error {
	if !hexRegex.MatchString(adrr) {
		return errors.New("invalid address")
	}

	return nil
}

func ValidateSignPayload(p *model.SigningPayload) error {
	if !hexRegex.MatchString(p.Address) {
		return errors.New("invalid address")
	}

	if !nonceRegex.MatchString(p.Nonce) {
		return errors.New("invalid nonce")
	}

	if len(p.Sig) == 0 {
		return errors.New("signature is missing")
	}

	return nil
}

func ValidateId(id string) error {
	return nil
}

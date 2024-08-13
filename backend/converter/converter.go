package converter

import (
	"errors"
	"regexp"
)

var hexRegex *regexp.Regexp = regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)

func ValidateAddress(adrr string) error {
	if !hexRegex.MatchString(adrr) {
		return errors.New("invalid address")
	}

	return nil
}

func ValidateId(id string) error {
	return nil
}

package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtHmacProvider struct {
	hmacSecret []byte
	issuer     string
	duration   time.Duration
}

func NewJwtHmacProvider(secret string, issuer string, duration time.Duration) *JwtHmacProvider {
	return &JwtHmacProvider{
		hmacSecret: []byte(secret),
		issuer:     issuer,
		duration:   duration,
	}
}

func (j *JwtHmacProvider) CreateStandart(s string) (string, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Issuer:    j.issuer,
		Subject:   s,
		ExpiresAt: jwt.NewNumericDate(now.Add(j.duration)),
		IssuedAt:  jwt.NewNumericDate(now),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.hmacSecret)
}

func (j *JwtHmacProvider) Verify(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return j.hmacSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("authentication error")
}

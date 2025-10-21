package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Subject string `json:"sub"`
	jwt.StandardClaims
}

func GenerateToken(secret, subject string, ttl time.Duration) (string, error) {
	claims := &Claims{
		Subject: subject,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

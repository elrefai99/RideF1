package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccess(payload string, secretToken string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": payload,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	})

	signedToken, err := token.SignedString([]byte(secretToken))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

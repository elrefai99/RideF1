package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenType string

const (
	AccessToken  TokenType = "access_token"
	RefreshToken TokenType = "refresh_token"
	PendingToken TokenType = "pending_token"
)

func GenerateTokens(payload any, tokenType TokenType) (string, error) {
	var secret string
	var lifetime time.Duration

	switch tokenType {
	case AccessToken:
		secret = os.Getenv("ACCESS_TOKEN")
		lifetime = 24 * time.Hour
	case RefreshToken:
		secret = os.Getenv("REFRESH_TOKEN")
		lifetime = 30 * 24 * time.Hour
	case PendingToken:
		secret = os.Getenv("PENDING_TOKEN")
		lifetime = 10 * time.Minute
	default:
		return "", errors.New("invalid token type")
	}

	if secret == "" {
		return "", errors.New("token signing secret is not configured")
	}

	now := time.Now()
	claims := jwt.MapClaims{
		"t_id": uuid.NewString(),
		"data": payload,
		"type": string(tokenType),
		"exp":  now.Add(lifetime).Unix(),
		"iat":  now.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	go func() {

	}()

	return signedToken, nil
}

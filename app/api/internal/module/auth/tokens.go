package auth

import (
	"context"
	"errors"
	"log"
	"os"
	"sync"
	"time"

	"github.com/elrefai99/RideF1/app/api/internal/Common/repository"
	"github.com/elrefai99/RideF1/app/api/internal/module/user"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TokenType string

const (
	AccessToken  TokenType = "access_token"
	RefreshToken TokenType = "refresh_token"
	PendingToken TokenType = "pending_token"
)

type IPayload struct {
	Uid      string
	User_id  uint
	Email    string
	Username string
}

func GenerateTokens(ctx context.Context, db *gorm.DB, payload IPayload, tokenType TokenType) (string, error) {
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

	if payload.User_id != 0 && tokenType == RefreshToken {
		body := user.Tokens{
			Uuid:         uuid.New(),
			User_id:      payload.User_id,
			RefrashToken: signedToken,
			ExpireAt:     now.Add(lifetime),
		}
		go func(ctx context.Context, db *gorm.DB, body user.Tokens) {
			var mux sync.RWMutex
			mux.Lock()

			defer mux.Unlock()

			var tokenData user.Tokens
			err := db.WithContext(ctx).Select("id").Where("user_id = ?", body.User_id).First(&tokenData).Error

			switch {
			case errors.Is(err, gorm.ErrRecordNotFound):
				createToken := repository.Repository[user.Tokens](db)
				tokens := []user.Tokens{body}
				if err := createToken.Create(ctx, &tokens); err != nil {
					log.Printf("could not save token: %v", err)
				}
			case err != nil:
				log.Printf("could not find token: %v", err)
			default:
				// A token already exists.
			}

		}(context.WithoutCancel(ctx), db, body)
	}

	return signedToken, nil
}

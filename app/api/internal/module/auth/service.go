package auth

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"
	"time"

	models "github.com/elrefai99/RideF1/Database/model"
	"github.com/elrefai99/RideF1/app/api/internal/Common/repository"
	"github.com/elrefai99/RideF1/app/api/internal/module/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IService struct {
	db     *gorm.DB
	schema *repository.IRepository[models.User]
}

func AuthService(db *gorm.DB) *IService {
	schema := repository.Repository[models.User](db)

	return &IService{
		db:     db,
		schema: schema,
	}
}

func (s *IService) RegisterService(ctx context.Context, payload user.CreateUserRequest) (bool, error) {

	query := map[string]any{
		"email":  payload.Email,
		"status": models.UserActive,
	}
	_, err := s.schema.FindOne(ctx, query, "id")

	if err == nil {
		return false, errors.New("Username or Email is already used")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}

	username := fmt.Sprintf(
		"%s_%d",
		strings.Join(strings.Fields(strings.ToLower(payload.Fullname)), "_"),
		rand.IntN(1000),
	)

	createUser := &models.User{
		Fullname:  payload.Fullname,
		Email:     strings.ToLower(payload.Email),
		Username:  username,
		Password:  string(hashPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	users := []models.User{*createUser}

	err = s.schema.Create(ctx, &users)
	if err != nil {
		return false, err
	}
	return true, nil
}

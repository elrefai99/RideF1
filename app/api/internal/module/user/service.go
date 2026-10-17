package user

import (
	"context"

	"gorm.io/gorm"
)

type IUserService struct {
	db *gorm.DB
}

func UserService(db *gorm.DB) *IUserService {
	return &IUserService{
		db: db,
	}
}

func (s *IUserService) UpdateUser(ctx context.Context) {
	// ctx, cancel := ctx.
}

func (s *IUserService) DeleteUser(ctx context.Context) {

}

func (s *IUserService) RiderProfile(ctx context.Context) {

}

func (s *IUserService) DriverProfile(ctx context.Context) {

}

package user

import (
	"context"

	"gorm.io/gorm"
)

type IUser struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) *IUser {
	return &IUser{
		db: db,
	}
}

func (r *IUser) FindById(ctx context.Context, id uint) {

}

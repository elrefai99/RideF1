package user

import (
	"context"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindById(ctx context.Context, id int64) (*User, error) {
	var user User

	err := r.db.WithContext(ctx).First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) FindByEmail(ctx context.Context, body User) (*User, error) {
	var user User

	err := r.db.WithContext(ctx).Where("email = ?", body.Email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

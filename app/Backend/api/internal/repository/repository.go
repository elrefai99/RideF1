package repository

import (
	"context"

	"gorm.io/gorm"
)

type IRepository[T any] struct {
	db *gorm.DB
}

type PayloadUser struct {
	id    uint
	email string
}

func Repository[T any](db *gorm.DB) *IRepository[T] {
	return &IRepository[T]{
		db: db,
	}
}

// Find
func (r *IRepository[T]) FindByID(ctx context.Context, id uint, col ...string) (*T, error) {
	var data T
	db := r.db.WithContext(ctx)

	if len(col) > 0 {
		db = db.Select(col)
	}
	err := db.First(&data, id).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *IRepository[T]) FindOne(ctx context.Context, payload map[string]any, col ...string) (*T, error) {
	var data T
	db := r.db.WithContext(ctx)

	if len(col) > 0 {
		db = db.Select(col)
	}

	err := db.Where(payload).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *IRepository[T]) Find(ctx context.Context, payload map[string]any, col ...string) (*[]T, error) {
	var data []T
	db := r.db.WithContext(ctx)

	if len(col) > 0 {
		db = db.Select(col)
	}

	err := db.Where(payload).Find(&data).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

// Create
func (r *IRepository[T]) Create(ctx context.Context, payload *[]T) error { // can work in create and create many
	return r.db.WithContext(ctx).Create(payload).Error
}

// Update
func (r *IRepository[T]) UpdateOne(ctx context.Context, query map[string]any, payload *[]T) error {
	return r.db.WithContext(ctx).Model(new(T)).Where(query).Updates(payload).Error
}
func (r *IRepository[T]) UpdateId(ctx context.Context, query map[string]any, payload *[]T) error {
	return r.db.WithContext(ctx).Model(new(T)).Where("id = ? ", query).Updates(payload).Error
}

// Delete
func (r *IRepository[T]) DeleteId(ctx context.Context, query map[string]any, payload *[]T) error {
	return r.db.WithContext(ctx).Model(new(T)).Where("id = ? ", query).Delete(payload).Error
}

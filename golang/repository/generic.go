package repository

import "gorm.io/gorm"

type GenericRepository[T any] struct {
	db *gorm.DB
}

func NewGenericRepository[T any](db *gorm.DB) *GenericRepository[T] {
	return &GenericRepository[T]{
		db: db,
	}
}

func (r *GenericRepository[T]) Create(data *T) (*T, error) {
	result := r.db.Create(data)
	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}

func (r *GenericRepository[T]) Delete(model T, id uint64) error {
	result := r.db.Delete(&model, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

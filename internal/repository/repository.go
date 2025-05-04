package repository

import "gorm.io/gorm"

type Repository[T any] interface {
	Create(entity *T) error
	Update(entity *T) error
	Delete(entity *T) error
	FindByField(field string, value any) ([]*T, error)
	FirstByField(field string, value any) (*T, error)
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) Repository[T] {
	return &baseRepository[T]{db: db}
}

func (r *baseRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *baseRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *baseRepository[T]) Delete(entity *T) error {
	return r.db.Delete(entity).Error
}

func (r *baseRepository[T]) FindByField(field string, value any) ([]*T, error) {
	var entities []*T
	if err := r.db.Where(field+" = ?", value).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *baseRepository[T]) FirstByField(field string, value any) (*T, error) {
	var entity *T
	if err := r.db.Where(field+" = ?", value).First(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

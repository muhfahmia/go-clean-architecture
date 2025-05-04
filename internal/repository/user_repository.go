package repository

import (
	"github.com/muhfahmia/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	// Add custom repository methods here
	Repository[entity.UserEntity]
}

type userRepository struct {
	db *gorm.DB
	Repository[entity.UserEntity]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		db:         db,
		Repository: NewBaseRepository[entity.UserEntity](db),
	}
}

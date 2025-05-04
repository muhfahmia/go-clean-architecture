package usecase

import (
	"github.com/muhfahmia/internal/repository"
)

type UserUsecase interface {
    Create() error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) Create() error {
	return nil
}

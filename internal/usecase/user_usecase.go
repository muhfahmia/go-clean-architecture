package usecase

import (
	"github.com/muhfahmia/internal/model"
	"github.com/muhfahmia/internal/repository"
	"github.com/muhfahmia/pkg/enum"
)

type UserUsecase interface {
	Create(req model.CreateUserRequest) model.AppError
}

type userUsecase struct {
	usecase BaseUsecase
	repo    repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository, usecase BaseUsecase) UserUsecase {
	return &userUsecase{usecase: usecase, repo: repo}
}

func (u *userUsecase) Create(req model.CreateUserRequest) model.AppError {
	if errV := u.usecase.Validate(req); errV != nil {
		return model.NewAppErrorValidation(errV, enum.ErrorValidation)
	}
	return nil
}

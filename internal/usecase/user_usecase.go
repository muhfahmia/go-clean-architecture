package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/muhfahmia/internal/model"
	"github.com/muhfahmia/internal/repository"
	"github.com/muhfahmia/pkg/enum"
)

type UserUsecase interface {
	Create(req model.CreateUserRequest) model.AppError
}

type userUsecase struct {
	validator *validator.Validate
	repo      repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository, validator *validator.Validate) UserUsecase {
	return &userUsecase{validator: validator, repo: repo}
}

func (u *userUsecase) Create(req model.CreateUserRequest) model.AppError {
	if errV := u.validator.Struct(req); errV != nil {
		return model.NewAppErrorValidation(errV, enum.ErrorValidation)
	}
	return nil
}

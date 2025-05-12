package usecase

import "github.com/go-playground/validator/v10"

type BaseUsecase interface {
	Validate(request any) error
}

type baseUsecase struct {
	validator *validator.Validate
}

func NewBaseUsecase(validator *validator.Validate) BaseUsecase {
	return &baseUsecase{
		validator: validator,
	}
}

func (b baseUsecase) Validate(request any) error {
	return b.validator.Struct(request)
}

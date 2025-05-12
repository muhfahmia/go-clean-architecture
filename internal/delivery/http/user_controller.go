package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfahmia/internal/model"
	"github.com/muhfahmia/internal/usecase"
	"github.com/muhfahmia/pkg/enum"
)

type UserController interface {
	Create(c *fiber.Ctx) error
}

type userControllerImpl struct {
	usecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) UserController {
	return &userControllerImpl{
		usecase: usecase,
	}
}

func (s *userControllerImpl) Create(c *fiber.Ctx) error {
	res := model.NewHttpResponseBuilder(c)
	req := model.CreateUserRequest{}

	if errP := res.WithRequestParameter(&req); errP != nil {
		return res.Send()
	}

	if err := s.usecase.Create(req); err != nil {
		res.WithError(err)
		return res.Send()
	}
	res.WithMessage(enum.MessageCreatedSuccess.Format("User"))
	res.WithHttpCode(fiber.StatusCreated)
	res.WithSuccess(true)
	return res.Send()
}

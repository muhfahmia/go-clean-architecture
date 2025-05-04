package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfahmia/internal/usecase"
)

type UserController interface {
	Create(c *fiber.Ctx) error
}

type userController struct {
	usecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) UserController {
	return &userController{
		usecase: usecase,
	}
}

func (s *userController) Create(c *fiber.Ctx) error {
	return nil
}

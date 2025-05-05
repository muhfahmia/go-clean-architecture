package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhfahmia/internal/model"
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
	res := model.NewAppResponse()
	req := model.CreateUserRequest{}
	if errB := c.BodyParser(&req); errB != nil {
	}

	if err := s.usecase.Create(req); err != nil {
		res.SetData("validationErrors", err.GetErrorValidation())
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}

package dependency

import (
	"github.com/muhfahmia/internal/delivery/http"
	"github.com/muhfahmia/internal/delivery/http/middleware"
)

type Container interface {
	GetAppMiddleware() middleware.Middleware
	GetUserController() http.UserController
}

type container struct {
	appMiddleware  middleware.Middleware
	userController http.UserController
}

func (a *appInjection) ProvideContainer() Container {
	container := container{}
	container.appMiddleware = a.NewAppMiddleware()

	userRepository := a.NewUserRepository()
	userUsecase := a.NewUserUsecase(userRepository)
	container.userController = http.NewUserController(userUsecase)
	return container
}

func (c container) GetAppMiddleware() middleware.Middleware {
	return c.appMiddleware
}

func (c container) GetUserController() http.UserController {
	return c.userController
}

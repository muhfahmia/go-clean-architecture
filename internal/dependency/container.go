package dependency

import (
	"github.com/muhfahmia/internal/config"
	"github.com/muhfahmia/internal/delivery/http"
	"github.com/muhfahmia/internal/delivery/http/middleware"
	"github.com/muhfahmia/internal/repository"
	"github.com/muhfahmia/internal/usecase"
)

type Container interface {
	GetAppMiddleware() middleware.Middleware
	GetUserController() http.UserController
}

type container struct {
	baseUsecase    usecase.BaseUsecase
	appMiddleware  middleware.Middleware
	userController http.UserController
}

func ProvideContainer(config config.AppConfig) Container {
	container := container{}
	container.appMiddleware = middleware.NewAppMiddleware()
	container.baseUsecase = usecase.NewBaseUsecase(config.GetValidator())

	userRepository := repository.NewUserRepository(config.GetPostgreSQLDatabase())
	userUsecase := usecase.NewUserUsecase(userRepository, container.baseUsecase)
	container.userController = http.NewUserController(userUsecase)
	return container
}

func (c container) GetAppMiddleware() middleware.Middleware {
	return c.appMiddleware
}

func (c container) GetUserController() http.UserController {
	return c.userController
}

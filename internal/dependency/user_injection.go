package dependency

import (
	"github.com/muhfahmia/internal/delivery/http"
	"github.com/muhfahmia/internal/repository"
	"github.com/muhfahmia/internal/usecase"
)

type UserInjection interface {
	NewUserController(usecase usecase.UserUsecase) http.UserController
	NewUserUsecase(repository repository.UserRepository) usecase.UserUsecase
	NewUserRepository() repository.UserRepository
}

func (di *appInjection) NewUserController(usecase usecase.UserUsecase) http.UserController {
	return http.NewUserController(usecase)
}

func (di *appInjection) NewUserUsecase(repository repository.UserRepository) usecase.UserUsecase {
	return usecase.NewUserUsecase(repository, di.config.GetValidator())
}

func (di *appInjection) NewUserRepository() repository.UserRepository {
	return repository.NewUserRepository(di.config.GetPostgreSQLDatabase())
}


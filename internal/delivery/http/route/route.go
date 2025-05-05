package route

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/muhfahmia/internal/dependency"
)

type Router interface {
	Setup()
}

type router struct {
	app       *fiber.App
	container dependency.Container
}

func NewRouter(container dependency.Container, app *fiber.App) Router {
	return router{
		app:       app,
		container: container,
	}
}

func (r router) Setup() {
	r.SetupGuest()
	r.SetupAuth()
}

func (r router) SetupGuest() {
	fmt.Println("Running Guest Route")
}

func (r router) SetupAuth() {
	r.app.Post("/auth/register", r.container.GetUserController().Create)
	fmt.Println("Running Auth Route")
}

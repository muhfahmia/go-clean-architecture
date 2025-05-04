package main

import (
	"github.com/muhfahmia/internal/config"
	"github.com/muhfahmia/internal/delivery/http/route"
	"github.com/muhfahmia/internal/dependency"
)

func main() {
	//setup config
	app := config.Bootstrap()

	//setup depedency injection
	var di dependency.Injection = dependency.NewAppInjection(app)
	container := di.ProvideContainer()

	//setup router
	router := route.NewRouter(container, app.GetApp())
	router.Setup()

	//running app
	app.Run()
}

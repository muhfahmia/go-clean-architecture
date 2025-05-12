package main

import (
	"github.com/muhfahmia/internal/config"
	"github.com/muhfahmia/internal/delivery/http/route"
	"github.com/muhfahmia/internal/dependency"
)

func main() {
	//setup config
	app := config.Bootstrap()
	container := dependency.ProvideContainer(app)

	//setup router
	router := route.NewRouter(container, app.GetApp())
	router.Setup()

	//running app
	app.Run()
}

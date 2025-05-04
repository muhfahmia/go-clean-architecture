package main

import (
	"github.com/muhfahmia/internal/config"
)

func main() {
	newConfig := config.NewAppConfig()
	config := newConfig.ProvideConfig()
	config.Run()
}

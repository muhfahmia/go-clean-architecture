package dependency

import (
	"github.com/muhfahmia/internal/config"
)

type Injection interface {
	ProvideContainer() Container
	MiddlewareInjection
	UserInjection
}

type appInjection struct {
	config config.AppConfig
}

func NewAppInjection(config config.AppConfig) Injection {
	return &appInjection{
		config: config,
	}
}

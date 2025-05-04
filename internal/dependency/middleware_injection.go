package dependency

import "github.com/muhfahmia/internal/delivery/http/middleware"

type MiddlewareInjection interface {
	NewAppMiddleware() middleware.Middleware
}

func (di *appInjection) NewAppMiddleware() middleware.Middleware {
	return middleware.NewAppMiddleware()
}

package middleware

import "github.com/gofiber/fiber/v2"

type Middleware interface {
	UserAuthMiddleware() fiber.Handler
}

type middleware struct {
	userAuth fiber.Handler
}

func NewAppMiddleware() Middleware {
	middleware := middleware{}
	middleware.userAuth = middleware.UserAuthMiddleware()
	return &middleware
}

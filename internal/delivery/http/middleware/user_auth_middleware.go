package middleware

import "github.com/gofiber/fiber/v2"

func (m *middleware) UserAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("X-User-Token")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "user_token_required",
				"message": "Invalid or missing authentication credentials",
			})
		}
		return c.Next()
	}
}

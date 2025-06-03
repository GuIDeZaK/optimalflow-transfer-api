package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/guide-backend/pkg/helpers"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Missing Authorization header")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid Authorization format")
		}

		token := parts[1]
		claims, err := helpers.ParseJWT(token)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token payload")
		}

		c.Locals("user_id", uint(userIDFloat))
		return c.Next()
	}
}

package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guide-backend/pkg/helpers"
)

func LogHelloMiddlewatre(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"data": "KUY"})
	}

	claims, err := helpers.ParseJWT(authHeader)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"data": "KUY"})
	}

	c.Locals("claims", claims)

	return c.Next()
}

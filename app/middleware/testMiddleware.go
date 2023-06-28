package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Test(c *fiber.Ctx) error {
	c.Locals("Test", "Hello GoFiber")
	return c.Next()
}

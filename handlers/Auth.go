package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func VerifyToken(c *fiber.Ctx) error {

	return c.Status(200).JSON("{}")
}

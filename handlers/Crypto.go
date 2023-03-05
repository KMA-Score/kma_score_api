package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"kma_score_api/utils"
	"kma_score_api/utils/aes"
)

func GenerateClientSecret(c *fiber.Ctx) error {
	key, _ := aes.GenerateAESKey()

	fmt.Println(key)

	return c.Status(200).JSON(utils.ApiResponse(200, "OK", key))
}

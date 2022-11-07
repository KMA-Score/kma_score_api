package utils

import "github.com/gofiber/fiber/v2"

func ApiResponse(statusCode int, message string, data any) fiber.Map {
	res := fiber.Map{
		"statusCode": statusCode,
		"message":    message,
		"data":       data,
	}

	return res
}

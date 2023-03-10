package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"kma_score_api/utils"
	"kma_score_api/utils/aes"
	"math/rand"
	"strings"
	"time"
)

// Credit: ChatGPT
func generateAPIKey() string {
	rand.Seed(time.Now().UnixNano())
	randomChar := string(rune(rand.Intn(26) + 97)) // lowercase alphabet

	// Get current date in ddmmyy format
	currentDate := time.Now().Format("020106")

	// Concatenate date and random character
	apiKey := fmt.Sprintf("%s%s", currentDate, randomChar)

	// Add additional random characters if necessary to reach length of 16
	for len(apiKey) < 24 {
		randChar := string(rune(rand.Intn(26) + 97))
		apiKey += randChar
	}

	// Prepend prefix to API key
	apiKey = "KMAS_" + apiKey

	// Convert key to uppercase
	apiKey = strings.ToUpper(apiKey)

	return apiKey
}

func GenerateClientSecret(c *fiber.Ctx) error {
	secretKey, err := aes.GenerateAESKey()

	if err != nil {
		return c.Status(500).JSON(utils.ApiResponse(500, "Generate key error", err))
	}

	apiKey := generateAPIKey()

	rsp := fiber.Map{
		"apiKey":    apiKey,
		"secretKey": secretKey,
	}

	return c.Status(200).JSON(utils.ApiResponse(200, "OK", rsp))
}

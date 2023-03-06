package middlewares

import (
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/utils"
	"kma_score_api/database"
	"kma_score_api/models"
	"kma_score_api/utils/aes"
	"log"
)

// Config defines the config for middleware.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// Enable api key system or not?
	//
	// Optional. Default: true
	Enable bool

	// Exclude domain list. There domains will not require key
	//
	// Optional. Default: []
	BackList []string

	// Api time deviation in second
	//
	// Optional. Default: 60
	ApiTimeDeviation int
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:             nil,
	Enable:           true,
	BackList:         []string{},
	ApiTimeDeviation: 60,
}

const (
	apiKeyField        = "X-KMA-API-KEY"
	apiSecretHashField = "X-KMA-API-SECRET-HASH"
)

// New creates a new middleware handler
func New(config ...Config) fiber.Handler {
	// Set default config
	cfg := ConfigDefault

	// Override config if provided
	if len(config) > 0 {
		cfg = config[0]

		// Set default values
		if cfg.Next == nil {
			cfg.Next = ConfigDefault.Next
		}
	}

	// Return new handler
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		// Don't execute if disabled
		if cfg.Enable == false {
			return c.Next()
		}

		// Don't execute if router is in black list
		routerPath := c.Route().Path

		if utils.Contains(cfg.BackList, routerPath) {
			return c.Next()
		}

		clientApiKey := c.Get(apiKeyField)
		clientSecretHash := c.Get(apiSecretHashField)

		if clientApiKey == "" || clientSecretHash == "" {
			return fiber.ErrForbidden
		}

		var key models.ApiKey
		result := database.DBConn.Model(&models.ApiKey{}).Where("`Key` = ?", clientApiKey).First(&key)

		if result.RowsAffected == 0 {
			log.Print(result.Error)
			return fiber.ErrForbidden
		}

		var keyDecoded, err = base64.StdEncoding.DecodeString(key.Secret)

		if err != nil {
			log.Print(result.Error)
			return fiber.ErrForbidden
		}

		var decoded string
		decoded, err = aes.DecryptGCM(keyDecoded, clientSecretHash)

		fmt.Println(decoded)

		//clientApiSecret := c.Get(apiSecretField)
		//
		//fmt.Println(reqHeaders)

		return c.Next()
	}
}

func ApiChecker() fiber.Handler {
	return New(Config{
		Enable:           true,
		BackList:         []string{},
		ApiTimeDeviation: 60000,
	})
}

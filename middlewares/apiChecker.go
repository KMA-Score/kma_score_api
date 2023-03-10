package middlewares

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/utils"
	"kma_score_api/database"
	"kma_score_api/models"
	kmaScoreUtils "kma_score_api/utils"
	"kma_score_api/utils/aes"
	"log"
	"os"
	"strconv"
	"time"
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
	Exclude []string

	// Api time deviation in second
	//
	// Optional. Default: 60
	ApiTimeDeviation int64
}

var DefaultConfig = Config{
	Next:             nil,
	Enable:           true,
	Exclude:          []string{},
	ApiTimeDeviation: 60,
}

const (
	apiKeyField        = "X-KMA-API-KEY"
	apiSecretHashField = "X-KMA-API-SECRET-HASH"
)

func New(config ...Config) fiber.Handler {
	// Set default config
	cfg := DefaultConfig

	// Override config if provided
	if len(config) > 0 {
		cfg = config[0]

		// Set default values
		if cfg.Next == nil {
			cfg.Next = DefaultConfig.Next
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
		routerPath := c.Path()

		if utils.Contains(cfg.Exclude, routerPath) {
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
			return c.Status(403).JSON(kmaScoreUtils.ApiResponse(403, "Key Not Found", nil))
		}

		var keyDecoded, err = base64.StdEncoding.DecodeString(key.Secret)

		if err != nil {
			log.Print("Base64 decode error: ", result.Error)
			return c.Status(500).JSON(kmaScoreUtils.ApiResponse(500, "Base64 decode error", nil))
		}

		var decoded []byte
		decoded, err = aes.DecryptCBC(keyDecoded, clientSecretHash)

		if err != nil {
			log.Print("Decrypt Error: ", err)
			return c.Status(500).JSON(kmaScoreUtils.ApiResponse(500, "Decrypt error or wrong key", nil))
		}

		//1678157885
		//e10974ffdc1a5d865c3c5e5e3f6c4f5d.e0ef8abe0a7f2c7a04b44958469b081c

		var currentTs int64
		currentTs, err = strconv.ParseInt(string(decoded), 10, 0)

		if err != nil {
			log.Print("ParseInt error: ", err)
			return c.Status(500).JSON(kmaScoreUtils.ApiResponse(500, "ParseInt error", nil))
		}

		currentTimestamp := time.Now().Unix()

		if currentTimestamp-currentTs > cfg.ApiTimeDeviation {
			log.Print("Oh No API expired")
			return c.Status(403).JSON(kmaScoreUtils.ApiResponse(403, "API Key Hash Expired", nil))
		}

		return c.Next()
	}
}

func ApiChecker() fiber.Handler {
	enableKey := false

	if enableKeyEnv, err := strconv.ParseBool(os.Getenv("ENABLE_API_KEY")); err == nil {
		enableKey = enableKeyEnv
	}

	return New(Config{
		Enable:           enableKey,
		Exclude:          []string{"/", "/api/aes/generateKey"},
		ApiTimeDeviation: 60, // 1 min
	})
}

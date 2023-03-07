package middlewares

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/utils"
	"kma_score_api/database"
	"kma_score_api/models"
	"kma_score_api/utils/aes"
	"log"
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
	BackList []string

	// Api time deviation in second
	//
	// Optional. Default: 60
	ApiTimeDeviation int64
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
		routerPath := c.Path()

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
			log.Print("Base64 decode error: ", result.Error)
			return fiber.ErrForbidden
		}

		var decoded []byte
		decoded, err = aes.DecryptCBC(keyDecoded, clientSecretHash)

		if err != nil {
			log.Print("Decrypt Error: ", result.Error)
			return fiber.ErrForbidden
		}

		//1678157885
		//e10974ffdc1a5d865c3c5e5e3f6c4f5d.e0ef8abe0a7f2c7a04b44958469b081c

		var currentTs int64
		currentTs, err = strconv.ParseInt(string(decoded), 10, 0)

		if err != nil {
			log.Print("ParseInt error: ", err)
			return fiber.ErrForbidden
		}

		currentTimestamp := time.Now().Unix()

		if currentTimestamp-currentTs > cfg.ApiTimeDeviation {
			log.Print("Oh No API expired")
			return fiber.ErrForbidden
		}

		return c.Next()
	}
}

func ApiChecker() fiber.Handler {
	return New(Config{
		Enable:           true,
		BackList:         []string{"/", "/api/aes/generateKey"},
		ApiTimeDeviation: 180, // 3 mins
	})
}

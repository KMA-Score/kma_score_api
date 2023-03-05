package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/utils"
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
	apiKeyField       = "X-KMA-API-KEY"
	apiSecretField    = "X-KMA-API-SECRET"
	apiTimestampField = "X-KMA-API-TIMESTAMP"
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

		//clientApiKey := c.Get(apiKeyField)
		//clientApiSecret := c.Get(apiSecretField)
		//clientApiTimestamp := c.Get(apiTimestampField)
		//
		//fmt.Println(reqHeaders)

		return c.Next()
	}
}

func ApiChecker() fiber.Handler {
	return New(Config{
		Enable:   true,
		BackList: []string{},
	})
}

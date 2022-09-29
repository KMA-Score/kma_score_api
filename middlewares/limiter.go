package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func Limiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        30,
		Expiration: 1 * time.Minute,
	})
}

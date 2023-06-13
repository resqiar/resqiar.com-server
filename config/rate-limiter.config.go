package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiterConfig() limiter.Config {
	config := limiter.Config{
		Max:        100,             // 100 request max
		Expiration: 1 * time.Minute, // 1 minutes max duration
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
	}

	return config
}

package config

import (
	"os"
	"resqiar.com-server/db"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var SessionStore *session.Store

func InitSession() {
	CLIENT_DOMAIN := os.Getenv("CLIENT_DOMAIN")

	SessionStore = session.New(session.Config{
		Expiration:     48 * time.Hour, // 2 days
		CookieHTTPOnly: true,
		CookieDomain:   CLIENT_DOMAIN,
		CookiePath:     "/",
		Storage:        db.RedisStore,
	})
}

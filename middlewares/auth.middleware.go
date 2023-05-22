package middlewares

import (
	"log"
	"resdev-server/config"

	"github.com/gofiber/fiber/v2"
)

func ProtectedRoute(c *fiber.Ctx) error {
	sess, err := config.SessionStore.Get(c)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	userID := sess.Get("ID")
	if userID == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// save user id from session into local key value
	c.Locals("userID", userID)

	return c.Next()
}

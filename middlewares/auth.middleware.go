package middlewares

import (
	"fmt"
	"log"

	"resqiar.com-server/config"

	"github.com/gofiber/fiber/v2"
)

func ProtectedRoute(c *fiber.Ctx) error {
	sess, err := config.SessionStore.Get(c)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	userID := sess.Get("ID")
	fmt.Println(userID)
	fmt.Println("======")
	fmt.Println(c.Cookies("session_id"))
	if userID == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// save user id from session into local key value
	c.Locals("userID", userID)

	return c.Next()
}

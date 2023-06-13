package middlewares

import (
	"resqiar.com-server/db"
	"resqiar.com-server/entities"

	"github.com/gofiber/fiber/v2"
)

// check logged in user if he/she is a tester,
// if so allow the route, else throw 401
func TesterRoute(c *fiber.Ctx) error {
	// user id from locals
	userID := c.Locals("userID")
	if userID == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	var currentUser entities.User
	result := db.DB.First(&currentUser, "ID = ? AND is_tester = ?", userID, true)

	// check if error OR if current user is NOT tester
	if result.Error != nil || !currentUser.IsTester {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}

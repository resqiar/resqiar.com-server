package handlers

import (
	"resdev-server/services"

	"github.com/gofiber/fiber/v2"
)

func SendUserProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	if userID == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	safeUser, err := services.FindUserByID(userID.(string))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	return c.JSON(&fiber.Map{
		"result": safeUser,
	})
}

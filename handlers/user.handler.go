package handlers

import (
	"fmt"

	"resqiar.com-server/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	SendUserProfile(c *fiber.Ctx) error
}

type UserHandlerImpl struct {
	UserService services.UserService
}

func (handler *UserHandlerImpl) SendUserProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	fmt.Println("SEND USER PROFILE")
	fmt.Println(userID)
	if userID == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	safeUser, err := handler.UserService.FindUserByID(userID.(string))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	return c.JSON(&fiber.Map{
		"result": safeUser,
	})
}

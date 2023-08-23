package handlers

import (
	"resqiar.com-server/inputs"
	"resqiar.com-server/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	SendCurrentUserProfile(c *fiber.Ctx) error
	SendUserProfile(c *fiber.Ctx) error
	SendUserUpdateProfile(c *fiber.Ctx) error
}

type UserHandlerImpl struct {
	UserService services.UserService
	UtilService services.UtilService
}

func (handler *UserHandlerImpl) SendCurrentUserProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID")
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

func (handler *UserHandlerImpl) SendUserProfile(c *fiber.Ctx) error {
	username := c.Params("username")
	if username == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	safeUser, err := handler.UserService.FindUserByUsername(username)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(&fiber.Map{
		"result": safeUser,
	})
}

func (handler *UserHandlerImpl) SendUserUpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	if userID == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	var payload inputs.UpdateUserInput
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := handler.UtilService.ValidateInput(payload); err != "" {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err,
		})
	}

	if payload.Username != "" {
		// check username if exist first, never ever proceed if isExist is true.
		isExist := handler.UserService.CheckUsernameExist(payload.Username)
		if isExist {
			return c.Status(fiber.StatusBadRequest).SendString("Username already exist")
		}
	}

	if err := handler.UserService.UpdateUser(&payload, userID.(string)); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

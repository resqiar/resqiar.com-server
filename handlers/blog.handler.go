package handlers

import (
	"resdev-server/inputs"
	"resdev-server/services"

	"github.com/gofiber/fiber/v2"
)

func SendBlogList(c *fiber.Ctx) error {
	result, err := services.GetAllBlogs()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"result": result,
	})
}

func SendBlogCreate(c *fiber.Ctx) error {
	// get current user ID
	userID := c.Locals("userID")
	if userID == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// define body payload
	var payload inputs.CreateBlogInput

	// bind the body parser into payload
	if err := c.BodyParser(&payload); err != nil {
		// send raw error (unprocessable entity)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// validate the payload using class-validator
	if err := services.ValidateInput(payload); err != "" {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err,
		})
	}

	result, err := services.CreateBlog(&payload, userID.(string))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"result": result,
	})
}

package handlers

import (
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

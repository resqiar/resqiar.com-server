package handlers

import (
	"resdev-server/inputs"
	"resdev-server/services"

	"github.com/gofiber/fiber/v2"
)

func SendBlogList(c *fiber.Ctx) error {
	result, err := services.GetAllBlogs(false)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"result": result,
	})
}

func SendPublishedBlogs(c *fiber.Ctx) error {
	// send only PUBLISHED blogs
	result, err := services.GetAllBlogs(true)
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
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"result": result,
	})
}

func SendCurrentUserBlogs(c *fiber.Ctx) error {
	userID := c.Locals("userID")

	result, err := services.GetCurrentUserBlogs(userID.(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"result": result,
	})
}

func SendPublishBlog(c *fiber.Ctx) error {
	// get current user ID
	userID := c.Locals("userID")

	// define body payload
	var payload inputs.PublishBlogInput

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

	_, err := services.ChangeBlogPublish(&payload, userID.(string), true)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}

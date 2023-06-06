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

func SendPublishedBlog(c *fiber.Ctx) error {
	// define body payload
	var payload inputs.BlogIDInput

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

	result, err := services.GetBlogDetail(payload.ID, true)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"result": result,
	})
}

func SendPublishedBlogs(c *fiber.Ctx) error {
	// send only PUBLISHED and SAFE blogs
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
	var payload inputs.BlogIDInput

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

func SendUnpublishBlog(c *fiber.Ctx) error {
	// get current user ID
	userID := c.Locals("userID")

	// define body payload
	var payload inputs.BlogIDInput

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

	_, err := services.ChangeBlogPublish(&payload, userID.(string), false)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}

func SendMyBlog(c *fiber.Ctx) error {
	// get current user ID
	userID := c.Locals("userID")

	// define body payload
	var payload inputs.BlogIDInput

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

	blog, err := services.GetBlogDetail(payload.ID, false)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	// if current requested blog is not the same author
	// as the one who request, return 404
	if blog.Author.ID != userID {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"result": blog,
	})
}

package handlers

import (
	"resqiar.com-server/constants"
	"resqiar.com-server/inputs"
	"resqiar.com-server/services"
	"resqiar.com-server/types"

	"github.com/gofiber/fiber/v2"
)

type BlogHandler interface {
	SendBlogList(c *fiber.Ctx) error
	SendPublishedBlog(c *fiber.Ctx) error
	SendPublishedBlogByID(c *fiber.Ctx) error
	SendPublishedBlogs(c *fiber.Ctx) error
	SendAuthorPublishedBlogs(c *fiber.Ctx) error
	SendPublishedSlugs(c *fiber.Ctx) error
	SendBlogCreate(c *fiber.Ctx) error
	SendCurrentUserBlogs(c *fiber.Ctx) error
	SendCurrentUserBlog(c *fiber.Ctx) error
	SendPublishBlog(c *fiber.Ctx) error
	SendUnpublishBlog(c *fiber.Ctx) error
	SendMyBlog(c *fiber.Ctx) error
	SendUpdateBlog(c *fiber.Ctx) error
}

type BlogHandlerImpl struct {
	BlogService services.BlogService
	UtilService services.UtilService
}

func (handler *BlogHandlerImpl) SendBlogList(c *fiber.Ctx) error {
	var qOrder string = c.Query("order", "DESC")

	// if order query does not exist in the map, set to default value
	if _, exist := constants.ValidOrders[qOrder]; !exist {
		qOrder = string(constants.DESC)
	}

	result, err := handler.BlogService.GetAllBlogs(false, constants.Order(qOrder))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"result": result,
	})
}

func (handler *BlogHandlerImpl) SendPublishedBlogByID(c *fiber.Ctx) error {
	ID := c.Params("id")

	result, err := handler.BlogService.GetBlogDetail(&types.BlogDetailOpts{
		GetBlogOpts: &types.GetBlogOpts{
			UseID:          ID,
			BlogSlug:       "",
			BlogAuthor:     "",
			IncludeContent: false,
			Published:      true,
		},
		ReturnHTML: false,
	})
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"result": result,
	})
}

func (handler *BlogHandlerImpl) SendPublishedBlog(c *fiber.Ctx) error {
	blogAuthor := c.Params("author")
	blogSlug := c.Params("slug")

	result, err := handler.BlogService.GetBlogDetail(&types.BlogDetailOpts{
		GetBlogOpts: &types.GetBlogOpts{
			UseID:          "",
			BlogAuthor:     blogAuthor,
			BlogSlug:       blogSlug,
			IncludeContent: true,
			Published:      true,
		},
		ReturnHTML: true,
	})

	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"result": result,
	})
}

func (handler *BlogHandlerImpl) SendPublishedBlogs(c *fiber.Ctx) error {
	var qOrder string = c.Query("order", "DESC")

	// if order query does not exist in the map, set to default value
	if _, exist := constants.ValidOrders[qOrder]; !exist {
		qOrder = string(constants.DESC)
	}

	// send only PUBLISHED and SAFE blogs
	result, err := handler.BlogService.GetAllBlogs(true, constants.Order(qOrder))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"result": result,
	})
}

func (handler *BlogHandlerImpl) SendAuthorPublishedBlogs(c *fiber.Ctx) error {
	var author string = c.Params("author")
	var qOrder string = c.Query("order", "DESC")

	// if order query does not exist in the map, set to default value
	if _, exist := constants.ValidOrders[qOrder]; !exist {
		qOrder = string(constants.DESC)
	}

	// send only PUBLISHED and SAFE blogs for specified author
	result, err := handler.BlogService.GetAllUserBlogs(author, constants.Order(qOrder))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"result": result,
	})
}

func (handler *BlogHandlerImpl) SendPublishedSlugs(c *fiber.Ctx) error {
	// send only PUBLISHED IDs
	result, err := handler.BlogService.GetAllSlugs()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"result": result,
	})
}

func (handler *BlogHandlerImpl) SendBlogCreate(c *fiber.Ctx) error {
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
	if err := handler.UtilService.ValidateInput(payload); err != "" {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err,
		})
	}

	result, err := handler.BlogService.CreateBlog(&payload, userID.(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"result": result,
	})
}

func (handler *BlogHandlerImpl) SendCurrentUserBlogs(c *fiber.Ctx) error {
	userID := c.Locals("userID")

	var qOrder string = c.Query("order", "DESC")

	// if order query does not exist in the map, set to default value
	if _, exist := constants.ValidOrders[qOrder]; !exist {
		qOrder = string(constants.DESC)
	}

	result, err := handler.BlogService.GetCurrentUserBlogs(userID.(string), constants.Order(qOrder))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"result": result,
	})
}

func (handler *BlogHandlerImpl) SendCurrentUserBlog(c *fiber.Ctx) error {
	userID := c.Locals("userID")

	// define body payload
	var payload inputs.BlogIDInput

	// bind the body parser into payload
	if err := c.BodyParser(&payload); err != nil {
		// send raw error (unprocessable entity)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// validate the payload using class-validator
	if err := handler.UtilService.ValidateInput(payload); err != "" {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err,
		})
	}

	result, err := handler.BlogService.GetCurrentUserBlog(payload.ID, userID.(string))
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"result": result,
	})
}

func (handler *BlogHandlerImpl) SendPublishBlog(c *fiber.Ctx) error {
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
	if err := handler.UtilService.ValidateInput(payload); err != "" {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err,
		})
	}

	err := handler.BlogService.ChangeBlogPublish(&payload, userID.(string), true)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (handler *BlogHandlerImpl) SendUnpublishBlog(c *fiber.Ctx) error {
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
	if err := handler.UtilService.ValidateInput(payload); err != "" {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err,
		})
	}

	err := handler.BlogService.ChangeBlogPublish(&payload, userID.(string), false)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (handler *BlogHandlerImpl) SendMyBlog(c *fiber.Ctx) error {
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
	if err := handler.UtilService.ValidateInput(payload); err != "" {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err,
		})
	}

	blog, err := handler.BlogService.GetBlogDetail(&types.BlogDetailOpts{
		GetBlogOpts: &types.GetBlogOpts{
			UseID:          payload.ID,
			BlogAuthor:     "",
			BlogSlug:       "",
			IncludeContent: true,
			Published:      false,
		},
		ReturnHTML: false,
	})
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

func (handler *BlogHandlerImpl) SendUpdateBlog(c *fiber.Ctx) error {
	// get current user ID
	userID := c.Locals("userID")

	// define body payload
	var payload inputs.UpdateBlogInput

	// bind the body parser into payload
	if err := c.BodyParser(&payload); err != nil {
		// send raw error (unprocessable entity)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// validate the payload using class-validator
	if err := handler.UtilService.ValidateInput(payload); err != "" {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err,
		})
	}

	err := handler.BlogService.EditBlog(&payload, userID.(string))
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendStatus(fiber.StatusOK)
}

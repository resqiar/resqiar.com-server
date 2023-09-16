package routes

import (
	"html/template"

	"github.com/gofiber/fiber/v2"
	"resqiar.com-server/constants"
	"resqiar.com-server/middlewares"
	"resqiar.com-server/services"
	"resqiar.com-server/types"
)

func InitWebRoute(server *fiber.App, userService *services.UserServiceImpl, blogService *services.BlogServiceImpl) {
	web := server.Group("/web")

	web.Get("/", middlewares.OptionalProtectedRoute, func(c *fiber.Ctx) error {
		ID := c.Locals("userID")
		if ID != nil {
			safeUser, _ := userService.FindUserByID(ID.(string))
			return c.Render("index", fiber.Map{
				"UserProfile":  safeUser,
				"ShowcaseData": constants.ShowcaseData,
				"Header":       0,
			})
		}

		return c.Render("index", fiber.Map{
			"UserProfile":  nil,
			"ShowcaseData": constants.ShowcaseData,
			"Header":       0,
		})
	})

	web.Get("/blog/:author/:slug", middlewares.OptionalProtectedRoute, func(c *fiber.Ctx) error {
		blogAuthor := c.Params("author")
		blogSlug := c.Params("slug")

		result, err := blogService.GetBlogDetail(&types.BlogDetailOpts{
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

		content := template.HTML(result.Content)

		ID := c.Locals("userID")
		if ID != nil {
			safeUser, _ := userService.FindUserByID(ID.(string))
			return c.Render("blog", fiber.Map{
				"UserProfile": safeUser,
				"Header":      1,
				"Blog":        result,
				"Content":     content,
			})
		}

		return c.Render("blog", fiber.Map{
			"UserProfile": nil,
			"Header":      1,
			"Blog":        result,
			"Content":     content,
		})
	})
}

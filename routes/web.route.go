package routes

import (
	"github.com/gofiber/fiber/v2"
	"resqiar.com-server/constants"
	"resqiar.com-server/middlewares"
	"resqiar.com-server/services"
)

func InitWebRoute(server *fiber.App, userService *services.UserServiceImpl) {
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
}

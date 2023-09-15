package routes

import (
	"github.com/gofiber/fiber/v2"
	"resqiar.com-server/constants"
)

func InitWebRoute(server *fiber.App) {
	web := server.Group("/web")

	web.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", constants.ShowcaseData)
	})
}

package routes

import (
	"resdev-server/handlers"
	"resdev-server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitAuthRoute(server *fiber.App) {
	auth := server.Group("/auth")

	auth.Get("/google", handlers.SendAuthGoogle)
	auth.Get("/google/callback", handlers.SendGoogleCallback)

	auth.Get("/logout", handlers.SendLogout)
	auth.Post("/status/adm",
		middlewares.ProtectedRoute,
		middlewares.AdminRoute,
		func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		},
	)

	// image-kit token based auth
	auth.Get("/ik", handlers.SendAuthIK)
}

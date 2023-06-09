package routes

import (
	"resdev-server/handlers"
	"resdev-server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitAuthRoute(server *fiber.App, handler handlers.AuthHandler) {
	auth := server.Group("/auth")

	auth.Get("/google", handler.SendAuthGoogle)
	auth.Get("/google/callback", handler.SendGoogleCallback)

	auth.Get("/logout", handler.SendLogout)
	auth.Post("/status/adm",
		middlewares.ProtectedRoute,
		middlewares.AdminRoute,
		func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		},
	)

	// image-kit token based auth
	auth.Get("/ik", handler.SendAuthIK)
}

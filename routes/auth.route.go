package routes

import (
	"resdev-server/handlers"

	"github.com/gofiber/fiber/v2"
)

func InitAuthRoute(server *fiber.App) {
	auth := server.Group("/auth")

	auth.Get("/google", handlers.SendAuthGoogle)
	auth.Get("/google/callback", handlers.SendGoogleCallback)
}

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
	auth.Get("/status", middlewares.ProtectedRoute, handlers.SendAuthStatus)
}

package routes

import (
	"resdev-server/handlers"

	"github.com/gofiber/fiber/v2"
)

func InitAuthRoute(server *fiber.App) {
	auth := server.Group("/auth")

	auth.Post("/register", handlers.SendHelloWord)
}

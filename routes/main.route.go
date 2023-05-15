package routes

import (
	"resdev-server/handlers"

	"github.com/gofiber/fiber/v2"
)

func InitMainRoutes(server *fiber.App) {
	server.Get("/", handlers.SendHelloWord)
}

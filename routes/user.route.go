package routes

import (
	"resdev-server/handlers"
	"resdev-server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitUserRoute(server *fiber.App) {
	user := server.Group("user")

	user.Get("/profile", middlewares.ProtectedRoute, handlers.SendUserProfile)
}

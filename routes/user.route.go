package routes

import (
	"resdev-server/handlers"
	"resdev-server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitUserRoute(server *fiber.App, handler handlers.UserHandler) {
	user := server.Group("user")

	user.Get("/profile", middlewares.ProtectedRoute, handler.SendUserProfile)
}

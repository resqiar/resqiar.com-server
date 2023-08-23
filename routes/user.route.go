package routes

import (
	"resqiar.com-server/handlers"
	"resqiar.com-server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitUserRoute(server *fiber.App, handler handlers.UserHandler) {
	user := server.Group("user")

	user.Get("/profile", middlewares.ProtectedRoute, handler.SendCurrentUserProfile)
	user.Get("/profile/:username", handler.SendUserProfile)
	// check username availability
	user.Get("/check/:username", middlewares.ProtectedRoute, handler.SendCheckUsername)

	user.Post("/profile/update", middlewares.ProtectedRoute, handler.SendUserUpdateProfile)
}

package routes

import (
	"resqiar.com-server/handlers"
	"resqiar.com-server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitUserRoute(server *fiber.App, handler handlers.UserHandler) {
	user := server.Group("user", middlewares.ProtectedRoute)

	user.Get("/profile", handler.SendUserProfile)

	user.Post("/profile/update", handler.SendUserUpdateProfile)
}

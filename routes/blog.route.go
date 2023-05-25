package routes

import (
	"resdev-server/handlers"
	"resdev-server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitBlogRoute(server *fiber.App) {
	blog := server.Group("/blog")

	blog.Get("/list", middlewares.ProtectedRoute, middlewares.AdminRoute, handlers.SendBlogList)
}

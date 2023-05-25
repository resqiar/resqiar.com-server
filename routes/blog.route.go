package routes

import (
	"resdev-server/handlers"
	"resdev-server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitBlogRoute(server *fiber.App) {
	blog := server.Group("/blog", middlewares.ProtectedRoute, middlewares.AdminRoute)

	blog.Get("/list", handlers.SendBlogList)
	blog.Post("/create", handlers.SendBlogCreate)
}

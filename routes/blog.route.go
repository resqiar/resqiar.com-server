package routes

import (
	"resdev-server/handlers"

	"github.com/gofiber/fiber/v2"
)

func InitBlogRoute(server *fiber.App) {
	blog := server.Group("/blog")

	blog.Get("/list", handlers.SendBlogList)
}

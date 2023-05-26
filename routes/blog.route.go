package routes

import (
	"resdev-server/handlers"
	"resdev-server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitBlogRoute(server *fiber.App) {
	blog := server.Group("/blog")

	// ONLY SEND PUBLISHED BLOGS FOR PUBLIC
	// drafted/unpublished blogs must only
	// be available to its author scope.
	blog.Get("/list", handlers.SendPublishedBlogs)

	// =========== SPECIAL ROUTES FOR ADM ONLY ===========
	blogADM := server.Group("/blog/admin", middlewares.ProtectedRoute, middlewares.AdminRoute)

	blogADM.Get("/list", handlers.SendBlogList)
	blogADM.Get("/list/current", handlers.SendCurrentUserBlogs)
	blogADM.Post("/create", handlers.SendBlogCreate)
}

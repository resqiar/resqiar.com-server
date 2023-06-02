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

	blog.Post("/create", middlewares.ProtectedRoute, middlewares.AdminRoute, handlers.SendBlogCreate)
	blog.Post("/publish", middlewares.ProtectedRoute, middlewares.AdminRoute, handlers.SendPublishBlog)
	blog.Get("/list/current", middlewares.ProtectedRoute, handlers.SendCurrentUserBlogs)

	// =========== SPECIAL ROUTES FOR ADM ONLY ===========
	blogADM := server.Group("/blog/adm", middlewares.ProtectedRoute, middlewares.AdminRoute)
	blogADM.Get("/list", handlers.SendBlogList)
}

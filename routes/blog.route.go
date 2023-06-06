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
	blog.Post("/get", handlers.SendPublishedBlog)

	blog.Get("/list/current", middlewares.ProtectedRoute, handlers.SendCurrentUserBlogs)
	blog.Post("/get/my", middlewares.ProtectedRoute, handlers.SendMyBlog)

	blog.Post("/create", middlewares.ProtectedRoute, middlewares.AdminRoute, handlers.SendBlogCreate)
	blog.Post("/publish", middlewares.ProtectedRoute, middlewares.AdminRoute, handlers.SendPublishBlog)
	blog.Post("/unpublish", middlewares.ProtectedRoute, middlewares.AdminRoute, handlers.SendUnpublishBlog)

	// =========== SPECIAL ROUTES FOR ADM ONLY ===========
	blogADM := server.Group("/blog/adm", middlewares.ProtectedRoute, middlewares.AdminRoute)
	blogADM.Get("/list", handlers.SendBlogList)
}

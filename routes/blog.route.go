package routes

import (
	"resqiar.com-server/handlers"
	"resqiar.com-server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitBlogRoute(server *fiber.App, handler handlers.BlogHandler) {
	blog := server.Group("/blog")

	// ONLY SEND PUBLISHED BLOGS FOR PUBLIC
	// drafted/unpublished blogs must only
	// be available to its author scope.
	blog.Get("/list", handler.SendPublishedBlogs)
	blog.Get("/list/slug", handler.SendPublishedSlugs)
	blog.Get("/get/published/:id", handler.SendPublishedBlogByID)
	blog.Get("/get/:author/:slug", handler.SendPublishedBlog)
	blog.Get("/get/:author", handler.SendAuthorPublishedBlogs)

	blog.Post("/list/current", middlewares.ProtectedRoute, handler.SendCurrentUserBlogs)
	blog.Post("/get/preview", middlewares.ProtectedRoute, handler.SendCurrentUserBlog)
	blog.Post("/get/my", middlewares.ProtectedRoute, handler.SendMyBlog)

	blog.Post("/create", middlewares.ProtectedRoute, handler.SendBlogCreate)
	blog.Post("/publish", middlewares.ProtectedRoute, handler.SendPublishBlog)
	blog.Post("/unpublish", middlewares.ProtectedRoute, handler.SendUnpublishBlog)
	blog.Post("/update", middlewares.ProtectedRoute, handler.SendUpdateBlog)

	// =========== SPECIAL ROUTES FOR ADM ONLY ===========
	blogADM := server.Group("/blog/adm", middlewares.ProtectedRoute, middlewares.AdminRoute)
	blogADM.Get("/list", handler.SendBlogList)
}

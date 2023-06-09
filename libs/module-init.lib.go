package libs

import (
	"resdev-server/handlers"
	"resdev-server/repositories"
	"resdev-server/routes"
	"resdev-server/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ModuleInit(server *fiber.App, DB *gorm.DB) {
	// Init repositories
	userRepository := repositories.InitUserRepo(DB)

	// Init services
	userService := services.UserServiceImpl{Repository: userRepository}
	utilService := services.UtilServiceImpl{}

	// Init handlers
	authHandler := handlers.AuthHandlerImpl{UserService: &userService}
	userHandler := handlers.UserHandlerImpl{UserService: &userService}
	blogHandler := handlers.BlogHandlerImpl{UtilService: &utilService}

	// Init routes
	routes.InitAuthRoute(server, &authHandler)
	routes.InitUserRoute(server, &userHandler)
	routes.InitBlogRoute(server, &blogHandler)
}

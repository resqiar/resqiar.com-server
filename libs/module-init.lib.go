package libs

import (
	"resqiar.com-server/handlers"
	"resqiar.com-server/repositories"
	"resqiar.com-server/routes"
	"resqiar.com-server/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ModuleInit(server *fiber.App, DB *gorm.DB) {
	// Init repositories
	userRepository := repositories.InitUserRepo(DB)
	blogRepository := repositories.InitBlogRepo(DB)

	// Init services
	userService := services.UserServiceImpl{Repository: userRepository}
	blogService := services.BlogServiceImpl{Repository: blogRepository}
	authService := services.AuthServiceImpl{}
	utilService := services.UtilServiceImpl{}

	// Init handlers
	authHandler := handlers.AuthHandlerImpl{
		UserService: &userService,
		AuthService: &authService,
	}
	userHandler := handlers.UserHandlerImpl{UserService: &userService}
	blogHandler := handlers.BlogHandlerImpl{
		BlogService: &blogService,
		UtilService: &utilService,
	}

	// Init routes
	routes.InitAuthRoute(server, &authHandler)
	routes.InitUserRoute(server, &userHandler)
	routes.InitBlogRoute(server, &blogHandler)
}

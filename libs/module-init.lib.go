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
	utilService := services.InitUtilService()
	userService := services.UserServiceImpl{
		Repository:  userRepository,
		UtilService: utilService,
	}
	blogService := services.BlogServiceImpl{UtilService: utilService, Repository: blogRepository}
	authService := services.AuthServiceImpl{}
	parserService := services.ParserServiceImpl{}

	// Init handlers
	authHandler := handlers.AuthHandlerImpl{
		UserService: &userService,
		AuthService: &authService,
		UtilService: utilService,
	}
	userHandler := handlers.UserHandlerImpl{
		UserService: &userService,
		UtilService: utilService,
	}
	blogHandler := handlers.BlogHandlerImpl{
		BlogService: &blogService,
		UtilService: utilService,
	}
	parserHandler := handlers.ParserHandlerImpl{
		ParserService: &parserService,
	}

	// Init routes
	routes.InitAuthRoute(server, &authHandler)
	routes.InitUserRoute(server, &userHandler)
	routes.InitBlogRoute(server, &blogHandler)
	routes.InitParserRoute(server, &parserHandler)
}

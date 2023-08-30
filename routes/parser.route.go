package routes

import (
	"resqiar.com-server/handlers"

	"github.com/gofiber/fiber/v2"
)

func InitParserRoute(server *fiber.App, handler handlers.ParserHandler) {
	parser := server.Group("/parser")
	parser.Post("/html", handler.ParseMDtoHTML)
}

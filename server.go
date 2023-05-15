package main

import (
	"os"
	"resdev-server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load env variables from .env file
	godotenv.Load()

	server := fiber.New()

	// Init routes
	routes.InitMainRoutes(server)

	PORT := os.Getenv("PORT")
	if err := server.Listen(":" + PORT); err != nil {
		panic(err)
	}
}

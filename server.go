package main

import (
	"os"
	"resdev-server/config"
	"resdev-server/db"
	"resdev-server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load env variables from .env file
	godotenv.Load()

	server := fiber.New()

	// Initialize database connection
	db.InitDB()

	// Initialize session
	config.InitSession()

	// Init routes
	routes.InitMainRoutes(server)
	routes.InitAuthRoute(server)

	PORT := os.Getenv("PORT")
	if err := server.Listen(":" + PORT); err != nil {
		panic(err)
	}
}

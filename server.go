package main

import (
	"os"
	"resdev-server/config"
	"resdev-server/db"
	"resdev-server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load env variables from .env file
	godotenv.Load()

	server := fiber.New()

	// Setup CORS
	server.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("CLIENT_URL"),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	db.InitDB()    // init Postgres db
	db.InitRedis() // init Redis db

	// Initialize session
	config.InitSession()

	// Init routes
	routes.InitAuthRoute(server)
	routes.InitBlogRoute(server)

	PORT := os.Getenv("PORT")
	if err := server.Listen(":" + PORT); err != nil {
		panic(err)
	}
}

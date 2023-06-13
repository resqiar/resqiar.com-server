package main

import (
	"os"
	"resqiar.com-server/config"
	"resqiar.com-server/db"
	"resqiar.com-server/libs"

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
		AllowOrigins:     os.Getenv("CLIENT_URL"),
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	DB := db.InitDB() // init Postgres db
	db.InitRedis()    // init Redis db

	// Initialize session
	config.InitSession()

	// Initialize repo, service, handler and route layers
	libs.ModuleInit(server, DB)

	PORT := os.Getenv("PORT")
	if err := server.Listen(":" + PORT); err != nil {
		panic(err)
	}
}

package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load env variables from .env file
	godotenv.Load()

	server := fiber.New()

	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	PORT := os.Getenv("PORT")
	if err := server.Listen(":" + PORT); err != nil {
		panic(err)
	}
}

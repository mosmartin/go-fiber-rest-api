package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mosmartin/go-fiber-rest-api/internal/handlers"
)

func healthCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{"status": "UP"})
}

func main() {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// create new app
	app := fiber.New()

	port := os.Getenv("PORT")

	app.Get("/api/v1/healthcheck", healthCheck)
	app.Post("/api/v1/products", handlers.CreateProduct)

	// start server
	log.Fatal(app.Listen(":" + port))
}

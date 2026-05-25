package main

import (
	"hotel-manager-api/internal/database"
	"hotel-manager-api/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Connect()

	routes.Setup(app)

	app.Listen(":3000")
}

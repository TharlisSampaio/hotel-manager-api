package routes

import (
	"hotel-manager-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/rooms", handlers.CreateRoom)
	app.Get("/rooms", handlers.GetRooms)
}

package handlers

import (
	"hotel-manager-api/internal/database"
	"hotel-manager-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

func CreateRoom(c *fiber.Ctx) error {
	room := new(models.Room)

	if err := c.BodyParser(room); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	database.DB.Create(&room)

	return c.Status(200).JSON(room)
}

func GetRooms(c *fiber.Ctx) error {
	rooms := []models.Room{}

	database.DB.Find(&rooms)

	return c.JSON(rooms)
}

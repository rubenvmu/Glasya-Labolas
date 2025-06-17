package controllers

import (
	"glasya-labolas/models"
	"github.com/gofiber/fiber/v2"
)

func GetEntities(c *fiber.Ctx) error {
	entities, err := models.ListEntities()
	if err != nil {
		return c.Status(500).SendString("Error cargando entidades")
	}
	return c.JSON(entities)
}
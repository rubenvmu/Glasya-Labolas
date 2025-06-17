package controllers

import (
	"glasya-labolas/models"
	"github.com/gofiber/fiber/v2"
)

func GetRelations(c *fiber.Ctx) error {
	relations, err := models.ListRelations()
	if err != nil {
		return c.Status(500).SendString("Error cargando relaciones")
	}
	return c.JSON(relations)
}
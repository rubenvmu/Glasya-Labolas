package controllers

import (
	"glasya-labolas/database"
	"github.com/gofiber/fiber/v2"
)

func GetStats(c *fiber.Ctx) error {
	db := database.GetDB()

	var entityCount int
	var relationCount int
	var lastLoad string

	db.QueryRow(`SELECT COUNT(*) FROM entities`).Scan(&entityCount)
	db.QueryRow(`SELECT COUNT(*) FROM relations`).Scan(&relationCount)
	db.QueryRow(`SELECT MAX(created_at) FROM entities`).Scan(&lastLoad)

	return c.JSON(fiber.Map{
		"entities":    entityCount,
		"relations":   relationCount,
		"last_loaded": lastLoad,
	})
}

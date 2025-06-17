package cmd

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"

	"glasya-labolas/controllers"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Inicia el servidor web glasya-labolas",
	Run: func(cmd *cobra.Command, args []string) {
		app := fiber.New()

		app.Get("/api/health", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"status":  "ok",
				"service": "glasya-labolas backend",
			})
		})

		app.Get("/api/entities", controllers.GetEntities)
		app.Get("/api/relations", controllers.GetRelations)
		app.Get("/api/stats", controllers.GetStats)

		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("🛰️ Glasya-Labolas API en línea.\nEndpoints: /api/health, /api/entities, /api/relations, /api/stats")
		})

		fmt.Println("🚀 Servidor disponible en http://localhost:8080")
		log.Fatal(app.Listen(":8080"))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

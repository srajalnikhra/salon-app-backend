package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/srajalnikhra/salon-app-backend/internal/config"
)

func main() {
	config.LoadEnv()

	appConfig := config.LoadAppConfig()

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Server is healthy",
			"data": fiber.Map{
				"app": appConfig.Name,
				"env": appConfig.Env,
			},
		})
	})

	log.Fatal(app.Listen(":" + appConfig.Port))
}

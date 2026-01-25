package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/srajalnikhra/salon-app-backend/internal/config"
	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
)

func main() {
	config.LoadEnv()

	appConfig := config.LoadAppConfig()
	dbConfig := config.LoadDBConfig()

	db.ConnectDatabase(dbConfig)

	err := db.DB.AutoMigrate(
		&models.Admin{},
		&models.Business{},
	)
	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Server is healthy",
			"data": fiber.Map{
				"app": appConfig.Name,
				"env": appConfig.Env,
				"db":  "connected",
			},
		})
	})

	log.Fatal(app.Listen(":" + appConfig.Port))
}

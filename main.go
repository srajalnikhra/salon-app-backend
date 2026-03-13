// @title Salon Management Backend API
// @version 1.0
// @description Multi-tenant Salon Management Backend built with Go + Fiber
// @termsOfService http://example.com/terms/

// @contact.name Srajal
// @contact.email srajal@example.com

// @host localhost:3000
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter: Bearer {your_jwt_token}

package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/gofiber/swagger"

	_ "github.com/srajalnikhra/salon-app-backend/docs"
	"github.com/srajalnikhra/salon-app-backend/internal/config"
	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/db/seed"
	"github.com/srajalnikhra/salon-app-backend/internal/routes"
	"github.com/srajalnikhra/salon-app-backend/internal/utils"
)

func main() {
	config.LoadEnv()

	appConfig := config.LoadAppConfig()
	dbConfig := config.LoadDBConfig()

	// Connect DB using GORM
	db.ConnectGorm(dbConfig)

	db.AutoMigrate()

	seed.Run()

	hash, _ := utils.HashPassword("1234")
	fmt.Println(hash)

	app := fiber.New()

	routes.RegisterRoutes(app)
	app.Get("/swagger/*", fiberSwagger.HandlerDefault)

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

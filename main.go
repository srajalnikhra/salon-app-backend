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
	"log"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/gofiber/swagger"

	_ "github.com/srajalnikhra/salon-app-backend/docs"
	"github.com/srajalnikhra/salon-app-backend/internal/config"
	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/db/seed"
	"github.com/srajalnikhra/salon-app-backend/internal/routes"
)

// main initializes and starts the Salon Management API server
func main() {
	// Load environment variables from .env file
	config.LoadEnv()

	// Load application configuration (name, env, port)
	appConfig := config.LoadAppConfig()
	// Load database configuration (host, port, user, password, etc.)
	dbConfig := config.LoadDBConfig()

	// Establish connection to PostgreSQL database
	db.ConnectGorm(dbConfig)

	// Auto-create/update database tables from model definitions
	db.AutoMigrate()

	// Populate database with initial sample data
	seed.Run()

	// Create new Fiber web framework instance
	app := fiber.New()

	// Register all API routes (public, protected, admin routes)
	routes.RegisterRoutes(app)
	// Setup Swagger API documentation endpoint
	app.Get("/swagger/*", fiberSwagger.HandlerDefault)

	// Health check endpoint for monitoring
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

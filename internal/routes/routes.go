package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/srajalnikhra/salon-app-backend/internal/controllers"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Post("/bookings", controllers.CreateBooking)
	api.Put("/bookings/:id/approve", controllers.ApproveBooking)
	api.Put("/bookings/:id/cancel", controllers.CancelBooking)
}

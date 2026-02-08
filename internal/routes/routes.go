package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/srajalnikhra/salon-app-backend/internal/controllers"
	"github.com/srajalnikhra/salon-app-backend/internal/middleware"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	// Auth
	api.Post("/admin/signup", controllers.AdminSignup)
	api.Post("/admin/login", controllers.AdminLogin)

	// Protected Booking APIs
	protected := api.Group("", middleware.AdminAuth())

	protected.Post("/bookings", controllers.CreateBooking)
	protected.Put("/bookings/:id/approve", controllers.ApproveBooking)
	protected.Put("/bookings/:id/cancel", controllers.CancelBooking)
	protected.Get("/bookings", controllers.ListBookings)
}

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
	api.Post("/staff/login", controllers.StaffLogin)

	// Protected (JWT required)
	protected := api.Group("", middleware.Auth())

	// Admin only routes
	admin := protected.Group("/admin", middleware.AdminOnly())
	admin.Post("/bookings", controllers.CreateBooking)
	admin.Put("/bookings/:id/approve", controllers.ApproveBooking)
	admin.Put("/bookings/:id/cancel", controllers.CancelBooking)

	admin.Post("/staff", controllers.CreateStaff)
	admin.Post("/staff/:staffId/services/:serviceId", controllers.AssignServiceToStaff)
	admin.Delete("/staff/:staffId/services/:serviceId", controllers.RemoveServiceFromStaff)

	// ✅ FIXED ENDPOINT (admin clearly owns this)
	admin.Post("/staff/:staffId/availability", controllers.SetStaffAvailability)

	// Staff + Admin
	protected.Get("/bookings", controllers.ListBookings)
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/srajalnikhra/salon-app-backend/internal/controllers"
	"github.com/srajalnikhra/salon-app-backend/internal/middleware"
)

// RegisterRoutes registers all API routes for the application
// Organizes routes into logical groups: public auth, protected, and admin-only
func RegisterRoutes(app *fiber.App) {
	// Create /api/v1 base route group for all endpoints
	api := app.Group("/api/v1")

	// ===== PUBLIC AUTH ROUTES (no JWT required) =====
	// Allow new admins to register and login
	api.Post("/admin/signup", controllers.AdminSignup)
	api.Post("/admin/login", controllers.AdminLogin)
	// Allow staff to login with credentials
	api.Post("/staff/login", controllers.StaffLogin)

	// ===== PROTECTED ROUTES (JWT required) =====
	// Apply Auth middleware to all protected routes
	// Middleware validates JWT token and extracts user info into request context
	protected := api.Group("", middleware.Auth())

	// ===== ADMIN-ONLY ROUTES (JWT + Admin role required) =====
	// Apply AdminOnly middleware on top of Auth middleware
	// Restricts access to admin users only
	admin := protected.Group("/admin", middleware.AdminOnly())
	// Booking management
	admin.Post("/bookings", controllers.CreateBooking)
	admin.Put("/bookings/:id/approve", controllers.ApproveBooking)
	admin.Put("/bookings/:id/cancel", controllers.CancelBooking)

	// Staff management
	admin.Post("/staff", controllers.CreateStaff)
	// Service-staff assignment
	admin.Post("/staff/:staffId/services/:serviceId", controllers.AssignServiceToStaff)
	admin.Delete("/staff/:staffId/services/:serviceId", controllers.RemoveServiceFromStaff)

	// Staff availability scheduling
	admin.Post("/staff/:staffId/availability", controllers.SetStaffAvailability)

	// ===== STAFF + ADMIN SHARED ROUTES =====
	// List bookings accessible to both staff and admins
	protected.Get("/bookings", controllers.ListBookings)

	// Service management
	admin.Post("/services", controllers.CreateService)
	protected.Get("/services", controllers.ListServices)
	protected.Get("/services/:id", controllers.GetServiceByID)
	admin.Put("/services/:id", controllers.UpdateService)
	admin.Delete("/services/:id", controllers.DeleteService)
}

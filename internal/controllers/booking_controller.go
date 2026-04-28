package controllers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/srajalnikhra/salon-app-backend/internal/controllers/dto"
	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
	"github.com/srajalnikhra/salon-app-backend/internal/services"
)

// CreateBooking handles booking creation with full validation
// Admin must provide customer info, service, and optional staff assignment
// Returns error if staff is unavailable or booking time conflicts
func CreateBooking(c *fiber.Ctx) error {
	// Parse and validate incoming booking request
	var req dto.CreateBookingRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	businessID := c.Locals("business_id").(uint)

	// Fetch service
	var service models.Service
	if err := db.DB.First(&service, req.ServiceID).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Service not found",
		})
	}

	// Calculate end time
	endTime := req.StartTime.Add(time.Duration(service.Duration) * time.Minute)

	// Check staff-service assignment
	if req.StaffID != nil {
		if !services.IsStaffAllowedForService(businessID, *req.StaffID, req.ServiceID) {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Staff is not assigned to this service",
			})
		}

		// Check if assigned staff is available during booking time
		if !services.IsStaffAvailableForBooking(businessID, *req.StaffID, req.StartTime, endTime) {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Staff is not available at the selected time",
			})
		}
	}

	// Check for scheduling conflicts (only if staff is assigned)
	if err := services.CheckBookingConflict(
		businessID,
		req.StaffID,
		req.StartTime,
		endTime,
	); err != nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// Find or create customer in database
	var customer models.Customer

	err := db.DB.
		Where("business_id = ? AND phone = ?", businessID, req.Customer.Phone).
		First(&customer).Error

	if err != nil {
		customer = models.Customer{
			BusinessID: businessID,
			Name:       req.Customer.Name,
			Phone:      req.Customer.Phone,
		}

		db.DB.Create(&customer)
	}

	// Create booking record with all details
	booking := models.Booking{
		BusinessID: businessID,
		CustomerID: customer.ID,
		ServiceID:  req.ServiceID,
		StaffID:    req.StaffID,
		Status:     "pending",
		StartTime:  req.StartTime,
		EndTime:    endTime,
		Notes:      req.Notes,
	}

	db.DB.Create(&booking)

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Booking created",
		"data": fiber.Map{
			"booking_id": booking.ID,
			"status":     booking.Status,
		},
	})
}

// ApproveBooking godoc
// @Summary Approve booking
// @Description Approve a pending booking (Admin only)
// @Tags Bookings
// @Produce json
// @Security BearerAuth
// @Param id path int true "Booking ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /admin/bookings/{id}/approve [put]
func ApproveBooking(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid booking ID",
		})
	}

	if err := services.ApproveBooking(uint(id)); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Booking approved",
	})
}

// CancelBooking godoc
// @Summary Cancel booking
// @Description Cancel a booking (Admin only)
// @Tags Bookings
// @Produce json
// @Security BearerAuth
// @Param id path int true "Booking ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /admin/bookings/{id}/cancel [put]
func CancelBooking(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid booking ID",
		})
	}

	if err := services.CancelBooking(uint(id)); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Booking cancelled",
	})
}

// ListBookings godoc
// @Summary List bookings
// @Description Get bookings (Admin or Staff)
// @Tags Bookings
// @Produce json
// @Security BearerAuth
// @Param date query string false "Filter by date (YYYY-MM-DD)"
// @Param staff_id query int false "Filter by staff ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /bookings [get]
func ListBookings(c *fiber.Ctx) error {
	// 1. Get business_id from JWT token (via middleware)
	businessID := c.Locals("business_id").(uint)

	// 2. optional date (YYYY-MM-DD)
	var date *time.Time
	if dateStr := c.Query("date"); dateStr != "" {
		parsed, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "invalid date format (use YYYY-MM-DD)",
			})
		}
		date = &parsed
	}

	// 3. optional staff_id
	var staffID *uint
	if sid := c.QueryInt("staff_id", 0); sid != 0 {
		id := uint(sid)
		staffID = &id
	}

	// 4. fetch bookings
	bookings, err := services.ListBookings(
		businessID,
		date,
		staffID,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "failed to fetch bookings",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    bookings,
	})
}

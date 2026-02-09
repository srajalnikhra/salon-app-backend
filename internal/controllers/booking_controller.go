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

func CreateBooking(c *fiber.Ctx) error {
	var req dto.CreateBookingRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	businessID := c.Locals("business_id").(uint)

	// 1. Fetch service (to get duration)
	var service models.Service
	if err := db.DB.First(&service, req.ServiceID).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Service not found",
		})
	}

	// 2. Calculate end time
	endTime := req.StartTime.Add(time.Duration(service.Duration) * time.Minute)

	// 3. Check staff-service assignment
	if req.StaffID != nil {
		if !services.IsStaffAllowedForService(businessID, *req.StaffID, req.ServiceID) {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Staff is not assigned to this service",
			})
		}
	}

	// 4. Conflict check
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

	// 4. Find or create customer
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

	// 5. Create booking
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

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

	// 3. Conflict check
	if err := services.CheckBookingConflict(
		req.BusinessID,
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
		Where("business_id = ? AND phone = ?", req.BusinessID, req.Customer.Phone).
		First(&customer).Error

	if err != nil {
		customer = models.Customer{
			BusinessID: req.BusinessID,
			Name:       req.Customer.Name,
			Phone:      req.Customer.Phone,
		}
		db.DB.Create(&customer)
	}

	// 5. Create booking
	booking := models.Booking{
		BusinessID: req.BusinessID,
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

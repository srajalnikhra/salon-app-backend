package controllers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/srajalnikhra/salon-app-backend/internal/services"
)

type SetStaffAvailabilityRequest struct {
	DayOfWeek int    `json:"day_of_week"`
	StartTime string `json:"start_time"` // "09:00"
	EndTime   string `json:"end_time"`   // "18:00"
}

func SetStaffAvailability(c *fiber.Ctx) error {
	businessID := c.Locals("business_id").(uint)
	staffID, err := c.ParamsInt("staffId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid staff ID",
		})
	}

	var req SetStaffAvailabilityRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	if req.DayOfWeek < 0 || req.DayOfWeek > 6 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid day_of_week (0-6)",
		})
	}

	// Parse times
	start, err := time.Parse("15:04", req.StartTime)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid start_time format (use HH:mm)",
		})
	}

	end, err := time.Parse("15:04", req.EndTime)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid end_time format (use HH:mm)",
		})
	}

	err = services.SetStaffAvailability(businessID, uint(staffID), req.DayOfWeek, start, end)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Staff availability set",
	})
}

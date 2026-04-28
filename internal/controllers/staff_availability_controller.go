package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
)

type SetStaffAvailabilityRequest struct {
	DayOfWeek int    `json:"day_of_week"`
	StartTime string `json:"start_time"` // "09:00"
	EndTime   string `json:"end_time"`   // "18:00"
}

// SetStaffAvailability godoc
// @Summary Set staff availability
// @Description Define working hours for a staff member
// @Tags Staff Availability
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param staffId path int true "Staff ID"
// @Param payload body dto.SetAvailabilityRequest true "Availability data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /admin/staff/{staffId}/availability [post]
func SetStaffAvailability(c *fiber.Ctx) error {

	staffID, err := c.ParamsInt("staffId")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "invalid staff id",
		})
	}

	// Extract business_id from JWT token
	businessID := c.Locals("business_id").(uint)

	// Parse request body
	var payload struct {
		DayOfWeek int    `json:"day_of_week"`
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "invalid request body",
		})
	}

	// Verify staff exists and belongs to this business
	var staff models.Staff
	if err := db.DB.Where("id = ? AND business_id = ?", staffID, businessID).
		First(&staff).Error; err != nil {

		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "staff not found",
		})
	}

	// Check if availability already exists for this day
	var existing models.StaffAvailability
	err = db.DB.Where(
		"staff_id = ? AND day_of_week = ?",
		staffID,
		payload.DayOfWeek,
	).First(&existing).Error

	if err == nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "availability already exists for this day",
		})
	}

	// Create availability record
	availability := models.StaffAvailability{
		BusinessID: businessID,
		StaffID:    uint(staffID),
		DayOfWeek:  payload.DayOfWeek,
		StartTime:  payload.StartTime,
		EndTime:    payload.EndTime,
	}

	db.DB.Create(&availability)

	// Return success response
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Staff availability set",
	})
}

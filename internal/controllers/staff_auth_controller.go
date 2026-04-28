package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/srajalnikhra/salon-app-backend/internal/controllers/dto"
	"github.com/srajalnikhra/salon-app-backend/internal/services"
	"github.com/srajalnikhra/salon-app-backend/internal/utils"
)

type StaffLoginRequest struct {
	Phone string `json:"phone"`
	PIN   string `json:"pin"`
}

// StaffLogin godoc
// @Summary Staff login
// @Description Login staff using phone and PIN
// @Tags Staff Auth
// @Accept json
// @Produce json
// @Param payload body StaffLoginRequest true "Staff login payload"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /staff/login [post]
func StaffLogin(c *fiber.Ctx) error {
	var req StaffLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false})
	}

	staff, err := services.StaffLogin(req.Phone, req.PIN)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	token, _ := utils.GenerateStaffToken(staff.ID, staff.BusinessID)

	return c.JSON(fiber.Map{
		"success": true,
		"token":   token,
	})
}

// CreateStaff godoc
// @Summary Create staff account
// @Description Create a new staff member (Admin only)
// @Tags Staff Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param payload body dto.CreateStaffRequest true "Staff creation payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Router /admin/staff [post]
func CreateStaff(c *fiber.Ctx) error {
	var req dto.CreateStaffRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	// Get business_id from JWT claims
	businessID := c.Locals("business_id").(uint)

	// Create staff
	staff, err := services.CreateStaff(businessID, req.Name, req.Phone, req.PIN, req.Role)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Staff created successfully",
		"data": fiber.Map{
			"id":          staff.ID,
			"name":        staff.Name,
			"phone":       staff.Phone,
			"role":        staff.Role,
			"is_active":   staff.IsActive,
			"business_id": staff.BusinessID,
		},
	})
}

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/srajalnikhra/salon-app-backend/internal/controllers/dto"
	"github.com/srajalnikhra/salon-app-backend/internal/services"
	"github.com/srajalnikhra/salon-app-backend/internal/utils"
)

// StaffLogin authenticates staff member and returns JWT token
// Uses phone and PIN instead of email/password
// StaffLogin godoc
// @Summary Staff login
// @Description Login as staff and receive JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.StaffLoginRequest true "Staff Login"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /staff/login [post]
func StaffLogin(c *fiber.Ctx) error {
	// Parse JSON request body
	var req dto.StaffLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false})
	}

	// Call service layer to validate staff credentials
	staff, err := services.StaffLogin(req.Phone, req.PIN)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// Generate JWT token for authenticated staff member
	token, _ := utils.GenerateStaffToken(staff.ID, staff.BusinessID)

	// Return success response with token
	return c.JSON(fiber.Map{
		"success": true,
		"token":   token,
	})
}

// CreateStaff allows admin to create new staff member
// Calls service layer for validation and DB operations
// CreateStaff godoc
// @Summary Create staff
// @Description Create a new staff member (Admin only)
// @Tags Staff
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateStaffRequest true "Staff Details"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /admin/staff [post]
func CreateStaff(c *fiber.Ctx) error {
	// Parse JSON request body into CreateStaffRequest DTO
	var req dto.CreateStaffRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	// Extract business_id from JWT claims set by Auth middleware
	businessID := c.Locals("business_id").(uint)

	// Call service layer to create staff record with validation
	staff, err := services.CreateStaff(businessID, req.Name, req.Phone, req.PIN, req.Role)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// Return success response with created staff details
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

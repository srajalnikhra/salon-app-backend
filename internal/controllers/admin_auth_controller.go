package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/srajalnikhra/salon-app-backend/internal/controllers/dto"
	"github.com/srajalnikhra/salon-app-backend/internal/services"
	"github.com/srajalnikhra/salon-app-backend/internal/utils"
)

// AdminSignup creates a new admin account and returns JWT token
// Calls service layer for business logic, then generates token for immediate login
// AdminSignup godoc
// @Summary Admin signup
// @Description Register a new admin account
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.AdminSignupRequest true "Admin Signup"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /admin/signup [post]
func AdminSignup(c *fiber.Ctx) error {
	// Parse JSON request body into AdminSignupRequest struct
	var req dto.AdminSignupRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false})
	}

	// Call service layer to validate and create admin account
	admin, err := services.AdminSignup(req.Name, req.Email, req.Password, req.BusinessID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	// Generate JWT token for newly created admin
	token, _ := utils.GenerateAdminToken(admin.ID, req.BusinessID)

	// Return success response with token
	return c.JSON(fiber.Map{
		"success": true,
		"token":   token,
	})
}

// AdminLogin authenticates admin and returns JWT token
// Validates email and password, then generates token on success
// AdminLogin godoc
// @Summary Admin login
// @Description Login as admin
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.AdminLoginRequest true "Admin Login"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /admin/login [post]
func AdminLogin(c *fiber.Ctx) error {
	// Parse JSON request body
	var req dto.AdminLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false})
	}

	// Call service layer to validate credentials
	admin, err := services.AdminLogin(req.Email, req.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	// Generate JWT token for authenticated admin
	token, _ := utils.GenerateAdminToken(admin.ID, 1)

	// Return success response with token
	return c.JSON(fiber.Map{
		"success": true,
		"token":   token,
	})
}

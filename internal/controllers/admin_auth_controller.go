package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/srajalnikhra/salon-app-backend/internal/services"
	"github.com/srajalnikhra/salon-app-backend/internal/utils"
)

// AdminSignupRequest holds required fields for admin account creation
type AdminSignupRequest struct {
	Name       string `json:"name"`        // Admin full name
	Email      string `json:"email"`       // Email for login (must be unique)
	Password   string `json:"password"`    // Password (will be hashed)
	BusinessID uint   `json:"business_id"` // Business/salon to manage
}

// AdminLoginRequest holds credentials for admin authentication
type AdminLoginRequest struct {
	Email    string `json:"email"`    // Registered email
	Password string `json:"password"` // Password
}

// AdminSignup creates a new admin account and returns JWT token
// Calls service layer for business logic, then generates token for immediate login
func AdminSignup(c *fiber.Ctx) error {
	// Parse JSON request body into AdminSignupRequest struct
	var req AdminSignupRequest
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
func AdminLogin(c *fiber.Ctx) error {
	// Parse JSON request body
	var req AdminLoginRequest
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

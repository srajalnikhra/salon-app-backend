package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/srajalnikhra/salon-app-backend/internal/services"
	"github.com/srajalnikhra/salon-app-backend/internal/utils"
)

type AdminSignupRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	BusinessID uint   `json:"business_id"`
}

type AdminLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AdminSignup godoc
// @Summary Admin signup
// @Description Create a new admin account and auto-login
// @Tags Admin Auth
// @Accept json
// @Produce json
// @Param payload body AdminSignupRequest true "Admin signup payload"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} ErrorResponse
// @Router /admin/signup [post]
func AdminSignup(c *fiber.Ctx) error {
	var req AdminSignupRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false})
	}

	admin, err := services.AdminSignup(req.Name, req.Email, req.Password, req.BusinessID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	token, _ := utils.GenerateAdminToken(admin.ID, req.BusinessID)

	return c.JSON(fiber.Map{
		"success": true,
		"token":   token,
	})
}

// AdminLogin godoc
// @Summary Admin login
// @Description Login admin using email and password
// @Tags Admin Auth
// @Accept json
// @Produce json
// @Param payload body AdminLoginRequest true "Admin login payload"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /admin/login [post]
func AdminLogin(c *fiber.Ctx) error {
	var req AdminLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false})
	}

	admin, err := services.AdminLogin(req.Email, req.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	// TEMP: business_id = 1 (later mapped properly)
	token, _ := utils.GenerateAdminToken(admin.ID, 1)

	return c.JSON(fiber.Map{
		"success": true,
		"token":   token,
	})
}
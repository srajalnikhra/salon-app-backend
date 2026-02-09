package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/srajalnikhra/salon-app-backend/internal/services"
	"github.com/srajalnikhra/salon-app-backend/internal/utils"
)

type StaffLoginRequest struct {
	Phone string `json:"phone"`
	PIN   string `json:"pin"`
}

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

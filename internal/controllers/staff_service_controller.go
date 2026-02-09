package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/srajalnikhra/salon-app-backend/internal/services"
)

func AssignServiceToStaff(c *fiber.Ctx) error {
	businessID := c.Locals("business_id").(uint)

	staffId, err := c.ParamsInt("staffId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid staff ID",
		})
	}

	serviceId, err := c.ParamsInt("serviceId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid service ID",
		})
	}

	if err := services.AssignServiceToStaff(businessID, uint(staffId), uint(serviceId)); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Service assigned to staff",
	})
}

func RemoveServiceFromStaff(c *fiber.Ctx) error {
	businessID := c.Locals("business_id").(uint)

	staffId, err := c.ParamsInt("staffId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid staff ID",
		})
	}

	serviceId, err := c.ParamsInt("serviceId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid service ID",
		})
	}

	if err := services.RemoveServiceFromStaff(businessID, uint(staffId), uint(serviceId)); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Service removed from staff",
	})
}

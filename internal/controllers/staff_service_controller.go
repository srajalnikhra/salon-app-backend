package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
	"github.com/srajalnikhra/salon-app-backend/internal/services"
)

// AssignServiceToStaff godoc
// @Summary Assign service to staff
// @Description Assign a service to a staff member (Admin only)
// @Tags Staff Services
// @Produce json
// @Security BearerAuth
// @Param staffId path int true "Staff ID"
// @Param serviceId path int true "Service ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /admin/staff/{staffId}/services/{serviceId} [post]
func AssignServiceToStaff(c *fiber.Ctx) error {

	staffID, err := c.ParamsInt("staffId")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "invalid staff id",
		})
	}

	serviceID, err := c.ParamsInt("serviceId")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "invalid service id",
		})
	}

	// Extract business_id from JWT token (tenant isolation)
	businessID := c.Locals("business_id").(uint)

	// Verify service exists
	var service models.Service
	if err := db.DB.First(&service, serviceID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "service not found",
		})
	}

	// Verify staff exists
	var staff models.Staff
	if err := db.DB.First(&staff, staffID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "staff not found",
		})
	}

	// Check if service is already assigned to this staff
	var existing models.StaffService
	err = db.DB.Where(
		"business_id = ? AND staff_id = ? AND service_id = ?",
		businessID,
		staffID,
		serviceID,
	).First(&existing).Error

	if err == nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "service already assigned to staff",
		})
	}

	assignment := models.StaffService{
		BusinessID: businessID,
		StaffID:    uint(staffID),
		ServiceID:  uint(serviceID),
	}

	db.DB.Create(&assignment)

	// Return success response
	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Service assigned to staff",
	})
}

// RemoveServiceFromStaff allows admin to unassign a service from staff member
// Service assignment is deleted from database
// RemoveServiceFromStaff godoc
// @Summary Remove service from staff
// @Description Remove an assigned service from a staff member (Admin only)
// @Tags Staff Services
// @Produce json
// @Security BearerAuth
// @Param staffId path int true "Staff ID"
// @Param serviceId path int true "Service ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /admin/staff/{staffId}/services/{serviceId} [delete]
func RemoveServiceFromStaff(c *fiber.Ctx) error {
	// Extract business_id from JWT token
	businessID := c.Locals("business_id").(uint)

	// Parse staff ID from URL path
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

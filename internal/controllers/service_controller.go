package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/srajalnikhra/salon-app-backend/internal/controllers/dto"
	"github.com/srajalnikhra/salon-app-backend/internal/services"
)

// CreateService godoc
// @Summary Create a new service
// @Description Creates a new salon service for the authenticated business (Admin only)
// @Tags Services
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateServiceRequest true "Service Details"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /admin/services [post]
func CreateService(c *fiber.Ctx) error {

	// Parse incoming request body
	var req dto.CreateServiceRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	// Extract business ID from JWT token
	businessID := c.Locals("business_id").(uint)

	// Call service layer
	service, err := services.CreateService(
		businessID,
		req.Name,
		req.Duration,
		req.Price,
	)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// Return success response
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Service created successfully",
		"data": fiber.Map{
			"id":          service.ID,
			"name":        service.Name,
			"duration":    service.Duration,
			"price":       service.Price,
			"is_active":   service.IsActive,
			"business_id": service.BusinessID,
		},
	})
}

// ListServices godoc
// @Summary List services
// @Description Returns all active services for the logged-in business
// @Tags Services
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /services [get]
func ListServices(c *fiber.Ctx) error {

	// Extract business_id from JWT token.
	// This ensures each business only sees its own services.
	businessID := c.Locals("business_id").(uint)

	// Fetch active services from service layer.
	services, err := services.ListServices(businessID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// Return list of services.
	return c.JSON(fiber.Map{
		"success": true,
		"data":    services,
	})
}

// GetServiceByID godoc
// @Summary Get service by ID
// @Description Returns a single active service belonging to the logged-in business
// @Tags Services
// @Produce json
// @Security BearerAuth
// @Param id path int true "Service ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /services/{id} [get]
func GetServiceByID(c *fiber.Ctx) error {

	// Parse service ID from URL.
	serviceID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid service ID",
		})
	}

	// Extract business ID from JWT.
	businessID := c.Locals("business_id").(uint)

	// Fetch service from service layer.
	service, err := services.GetServiceByID(businessID, uint(serviceID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// Return service details.
	return c.JSON(fiber.Map{
		"success": true,
		"data":    service,
	})
}

// UpdateService godoc
// @Summary Update service
// @Description Updates an existing service for the authenticated business (Admin only)
// @Tags Services
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Service ID"
// @Param request body dto.UpdateServiceRequest true "Updated Service"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /admin/services/{id} [put]
func UpdateService(c *fiber.Ctx) error {

	// Parse service ID from URL.
	serviceID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid service ID",
		})
	}

	// Parse request body.
	var req dto.UpdateServiceRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	// Extract business ID from JWT.
	businessID := c.Locals("business_id").(uint)

	// Call service layer.
	service, err := services.UpdateService(
		businessID,
		uint(serviceID),
		req.Name,
		req.Duration,
		req.Price,
	)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// Return updated service.
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Service updated successfully",
		"data": fiber.Map{
			"id":       service.ID,
			"name":     service.Name,
			"duration": service.Duration,
			"price":    service.Price,
		},
	})
}

// DeleteService godoc
// @Summary Delete service
// @Description Soft deletes a service by marking it inactive (Admin only)
// @Tags Services
// @Produce json
// @Security BearerAuth
// @Param id path int true "Service ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /admin/services/{id} [delete]
func DeleteService(c *fiber.Ctx) error {

	// Parse service ID from URL.
	serviceID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid service ID",
		})
	}

	// Extract business ID from JWT.
	businessID := c.Locals("business_id").(uint)

	// Call service layer.
	if err := services.DeleteService(
		businessID,
		uint(serviceID),
	); err != nil {

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Service deleted successfully",
	})
}
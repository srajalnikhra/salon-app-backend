package services

import (
	"errors"
	"strings"

	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
)

// CreateService creates a new service for a business
// Validates required fields and prevents duplicate service names
func CreateService(
	businessID uint,
	name string,
	duration int,
	price float64,
) (*models.Service, error) {

	// Validate required fields
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("service name is required")
	}

	if duration <= 0 {
		return nil, errors.New("service duration must be greater than 0")
	}

	if price <= 0 {
		return nil, errors.New("service price must be greater than 0")
	}

	// Check if service name already exists for this business (case-insensitive)
	var existing models.Service

	err := db.DB.
		Where(
			"business_id = ? AND LOWER(name) = LOWER(?)",
			businessID,
			strings.TrimSpace(name),
		).
		First(&existing).Error

	if err == nil {
		return nil, errors.New("service already exists for this business")
	}

	// Create new service
	service := models.Service{
		BusinessID: businessID,
		Name:       strings.TrimSpace(name),
		Duration:   duration,
		Price:      price,
		IsActive:   true,
	}

	// Save service
	if err := db.DB.Create(&service).Error; err != nil {
		return nil, errors.New("failed to create service")
	}

	return &service, nil
}

// ServiceListItem represents the response returned to clients.
// It exposes only the fields required by the frontend.
type ServiceListItem struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Duration int     `json:"duration"`
	Price    float64 `json:"price"`
}

// ListServices returns all active services for a business.
// Only minimal fields are selected to keep the response lightweight.
func ListServices(businessID uint) ([]ServiceListItem, error) {

	// Slice to store the services returned from the database.
	var services []ServiceListItem

	// Fetch only active services belonging to this business.
	err := db.DB.
		Model(&models.Service{}).
		Select("id", "name", "duration", "price").
		Where("business_id = ? AND is_active = ?", businessID, true).
		Order("name ASC").
		Find(&services).Error

	if err != nil {
		return nil, errors.New("failed to fetch services")
	}

	return services, nil
}

// GetServiceByID returns a single active service for the given business.
// It ensures businesses cannot access each other's services.
func GetServiceByID(
	businessID uint,
	serviceID uint,
) (*ServiceListItem, error) {

	// Store fetched service.
	var service ServiceListItem

	// Fetch service belonging to the logged-in business.
	err := db.DB.
		Model(&models.Service{}).
		Select("id", "name", "duration", "price").
		Where(
			"business_id = ? AND id = ? AND is_active = ?",
			businessID,
			serviceID,
			true,
		).
		First(&service).Error

	if err != nil {
		return nil, errors.New("service not found")
	}

	return &service, nil
}

// UpdateService updates an existing service for the logged-in business.
// It validates input, prevents duplicate service names, and ensures
// businesses cannot modify services belonging to another business.
func UpdateService(
	businessID uint,
	serviceID uint,
	name string,
	duration int,
	price float64,
) (*models.Service, error) {

	// Validate service name.
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("service name is required")
	}

	// Validate duration.
	if duration <= 0 {
		return nil, errors.New("service duration must be greater than 0")
	}

	// Validate price.
	if price <= 0 {
		return nil, errors.New("service price must be greater than 0")
	}

	// Find the service that belongs to this business.
	var service models.Service

	err := db.DB.
		Where(
			"business_id = ? AND id = ? AND is_active = ?",
			businessID,
			serviceID,
			true,
		).
		First(&service).Error

	if err != nil {
		return nil, errors.New("service not found")
	}

	// Prevent duplicate service names (excluding the current service).
	var existing models.Service

	err = db.DB.
		Where(
			"business_id = ? AND LOWER(name) = LOWER(?) AND id <> ?",
			businessID,
			strings.TrimSpace(name),
			serviceID,
		).
		First(&existing).Error

	if err == nil {
		return nil, errors.New("service with this name already exists")
	}

	// Update service fields.
	service.Name = strings.TrimSpace(name)
	service.Duration = duration
	service.Price = price

	// Save changes.
	if err := db.DB.Save(&service).Error; err != nil {
		return nil, errors.New("failed to update service")
	}

	return &service, nil
}

// DeleteService performs a soft delete by marking the service as inactive.
// This preserves historical booking data while hiding the service
// from future API responses.
func DeleteService(
	businessID uint,
	serviceID uint,
) error {

	// Find the service belonging to the logged-in business.
	var service models.Service

	err := db.DB.
		Where(
			"business_id = ? AND id = ? AND is_active = ?",
			businessID,
			serviceID,
			true,
		).
		First(&service).Error

	if err != nil {
		return errors.New("service not found")
	}

	// Soft delete.
	service.IsActive = false

	// Save updated status.
	if err := db.DB.Save(&service).Error; err != nil {
		return errors.New("failed to delete service")
	}

	return nil
}
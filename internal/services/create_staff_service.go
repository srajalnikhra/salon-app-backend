package services

import (
	"errors"

	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
	"github.com/srajalnikhra/salon-app-backend/internal/utils"
)

// CreateStaff creates new staff member account
// Validates required fields, checks phone uniqueness, and hashes PIN
func CreateStaff(businessID uint, name, phone, pin, role string) (*models.Staff, error) {
	// Validate that all required fields are provided
	if name == "" || phone == "" || pin == "" || role == "" {
		return nil, errors.New("missing required fields")
	}

	// Check if phone number already exists for this business
	var existing models.Staff
	if err := db.DB.Where("business_id = ? AND phone = ?", businessID, phone).First(&existing).Error; err == nil {
		return nil, errors.New("phone number already exists for this business")
	}

	// Hash PIN for secure storage
	hashedPIN, _ := utils.HashPassword(pin)

	// Create staff record
	staff := models.Staff{
		BusinessID: businessID,
		Name:       name,
		Phone:      phone,
		PIN:        hashedPIN,
		Role:       role,
		IsActive:   true,
	}

	if err := db.DB.Create(&staff).Error; err != nil {
		return nil, errors.New("failed to create staff")
	}

	return &staff, nil
}

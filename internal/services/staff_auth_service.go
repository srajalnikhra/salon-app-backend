package services

import (
	"errors"

	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
	"github.com/srajalnikhra/salon-app-backend/internal/utils"
)

// StaffLogin authenticates staff member by phone and PIN
// Returns staff details if credentials are valid
func StaffLogin(phone, pin string) (*models.Staff, error) {
	// Find staff by phone number
	var staff models.Staff

	if err := db.DB.Where("phone = ?", phone).First(&staff).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Check if staff is active
	if !staff.IsActive {
		return nil, errors.New("staff inactive")
	}

	// Validate PIN against hashed PIN
	if err := utils.CheckPassword(staff.PIN, pin); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &staff, nil
}

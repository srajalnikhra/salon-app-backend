package services

import (
	"errors"

	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
	"github.com/srajalnikhra/salon-app-backend/internal/utils"
)

func StaffLogin(phone, pin string) (*models.Staff, error) {
	var staff models.Staff

	if err := db.DB.Where("phone = ?", phone).First(&staff).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !staff.IsActive {
		return nil, errors.New("staff inactive")
	}

	if err := utils.CheckPassword(staff.PIN, pin); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &staff, nil
}

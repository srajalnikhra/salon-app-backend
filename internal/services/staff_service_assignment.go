package services

import (
	"errors"

	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
)

func AssignServiceToStaff(businessID, staffID, serviceID uint) error {
	// Check if already assigned
	var count int64
	db.DB.Model(&models.StaffService{}).
		Where("business_id = ? AND staff_id = ? AND service_id = ?", businessID, staffID, serviceID).
		Count(&count)

	if count > 0 {
		return errors.New("service already assigned to staff")
	}

	assignment := models.StaffService{
		BusinessID: businessID,
		StaffID:    staffID,
		ServiceID:  serviceID,
	}

	return db.DB.Create(&assignment).Error
}

func RemoveServiceFromStaff(businessID, staffID, serviceID uint) error {
	result := db.DB.Where("business_id = ? AND staff_id = ? AND service_id = ?", businessID, staffID, serviceID).
		Delete(&models.StaffService{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("assignment not found")
	}
	return nil
}

func IsStaffAllowedForService(businessID, staffID, serviceID uint) bool {
	var count int64
	db.DB.Model(&models.StaffService{}).
		Where("business_id = ? AND staff_id = ? AND service_id = ?", businessID, staffID, serviceID).
		Count(&count)

	return count > 0
}

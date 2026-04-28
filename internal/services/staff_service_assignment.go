package services

import (
	"errors"

	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
)

// AssignServiceToStaff creates a staff-service relationship
// Prevents duplicate assignments
func AssignServiceToStaff(businessID, staffID, serviceID uint) error {
	// Check if already assigned
	var count int64
	db.DB.Model(&models.StaffService{}).
		Where("business_id = ? AND staff_id = ? AND service_id = ?", businessID, staffID, serviceID).
		Count(&count)

	if count > 0 {
		return errors.New("service already assigned to staff")
	}

	// Create assignment record
	assignment := models.StaffService{
		BusinessID: businessID,
		StaffID:    staffID,
		ServiceID:  serviceID,
	}

	return db.DB.Create(&assignment).Error
}

// RemoveServiceFromStaff deletes a staff-service assignment
func RemoveServiceFromStaff(businessID, staffID, serviceID uint) error {
	// Delete the assignment record
	result := db.DB.Where("business_id = ? AND staff_id = ? AND service_id = ?", businessID, staffID, serviceID).
		Delete(&models.StaffService{})

	if result.Error != nil {
		return result.Error
	}
	// Ensure at least one record was deleted
	if result.RowsAffected == 0 {
		return errors.New("assignment not found")
	}
	return nil
}

// IsStaffAllowedForService checks if staff is qualified to offer a service
func IsStaffAllowedForService(businessID, staffID, serviceID uint) bool {
	// Count assignment records for this staff-service combination
	var count int64
	db.DB.Model(&models.StaffService{}).
		Where("business_id = ? AND staff_id = ? AND service_id = ?", businessID, staffID, serviceID).
		Count(&count)

	// Return true if assignment exists
	return count > 0
}

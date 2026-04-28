package services

import (
	"errors"
	"time"

	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
)

var (
	ErrBookingConflict     = errors.New("booking time conflicts with an existing booking")
	ErrBookingNotFound     = errors.New("booking not found")
	ErrInvalidBookingState = errors.New("invalid booking state")
)

// CheckBookingConflict verifies that booking time doesn't overlap with existing bookings
// If staff is assigned, checks only that staff's bookings; otherwise checks all for that service
func CheckBookingConflict(
	businessID uint,
	staffID *uint,
	startTime time.Time,
	endTime time.Time,
) error {
	// Query for overlapping bookings (active or pending)
	query := db.DB.Model(&models.Booking{}).
		Where("business_id = ?", businessID).
		Where("status IN ?", []string{"pending", "confirmed"}).
		Where("start_time < ? AND end_time > ?", endTime, startTime)

	// If staff is assigned, check conflict for that staff
	if staffID != nil {
		query = query.Where("staff_id = ?", *staffID)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return err
	}

	// If any conflicting bookings found, return error
	if count > 0 {
		return ErrBookingConflict
	}

	return nil
}

// ApproveBooking changes booking status from pending to confirmed
func ApproveBooking(bookingID uint) error {
	// Fetch booking from database
	var booking models.Booking

	if err := db.DB.First(&booking, bookingID).Error; err != nil {
		return ErrBookingNotFound
	}

	// Validate booking is in pending state
	if booking.Status != "pending" {
		return ErrInvalidBookingState
	}

	// Update status to confirmed
	booking.Status = "confirmed"
	return db.DB.Save(&booking).Error
}

// CancelBooking changes booking status to cancelled
func CancelBooking(bookingID uint) error {
	// Fetch booking from database
	var booking models.Booking

	if err := db.DB.First(&booking, bookingID).Error; err != nil {
		return ErrBookingNotFound
	}

	// Ensure booking is not already cancelled
	if booking.Status == "cancelled" {
		return ErrInvalidBookingState
	}

	// Update status to cancelled
	booking.Status = "cancelled"
	return db.DB.Save(&booking).Error
}

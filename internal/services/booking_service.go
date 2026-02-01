package services

import (
	"errors"
	"time"

	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
)

var (
	ErrBookingConflict       = errors.New("booking time conflicts with an existing booking")
	ErrBookingNotFound       = errors.New("booking not found")
	ErrInvalidBookingState   = errors.New("invalid booking state")
)

func CheckBookingConflict(
	businessID uint,
	staffID *uint,
	startTime time.Time,
	endTime time.Time,
) error {

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

	if count > 0 {
		return ErrBookingConflict
	}

	return nil
}


// Approve booking
func ApproveBooking(bookingID uint) error {
	var booking models.Booking

	if err := db.DB.First(&booking, bookingID).Error; err != nil {
		return ErrBookingNotFound
	}

	if booking.Status != "pending" {
		return ErrInvalidBookingState
	}

	booking.Status = "confirmed"
	return db.DB.Save(&booking).Error
}

// Cancel booking
func CancelBooking(bookingID uint) error {
	var booking models.Booking

	if err := db.DB.First(&booking, bookingID).Error; err != nil {
		return ErrBookingNotFound
	}

	if booking.Status == "cancelled" {
		return ErrInvalidBookingState
	}

	booking.Status = "cancelled"
	return db.DB.Save(&booking).Error
}

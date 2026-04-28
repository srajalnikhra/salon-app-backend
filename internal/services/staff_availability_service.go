package services

import (
	"errors"
	"time"

	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
)

// SetStaffAvailability creates working hours schedule for staff on a specific day
// Prevents duplicate entries for same day and validates time logic
func SetStaffAvailability(businessID, staffID uint, dayOfWeek int, start, end time.Time) error {
	// Extract just HH:mm from time values
	startStr := start.Format("15:04")
	endStr := end.Format("15:04")

	// Validate start time is before end time
	if startStr >= endStr {
		return errors.New("start time must be before end time")
	}

	// Prevent duplicate day entries
	var count int64
	err := db.DB.Model(&models.StaffAvailability{}).
		Where("business_id = ? AND staff_id = ? AND day_of_week = ?", businessID, staffID, dayOfWeek).
		Count(&count).Error

	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("availability already exists for this day")
	}

	availability := models.StaffAvailability{
		BusinessID: businessID,
		StaffID:    staffID,
		DayOfWeek:  dayOfWeek,
		StartTime:  startStr,
		EndTime:    endStr,
		CreatedAt:  time.Now(),
	}

	return db.DB.Create(&availability).Error
}

// IsStaffAvailableForBooking checks if staff is available during requested booking time
// Validates day of week and time overlap with staff availability
func IsStaffAvailableForBooking(businessID, staffID uint, bookingStart, bookingEnd time.Time) bool {
	// Determine day of week (0=Sunday, 6=Saturday)
	dayOfWeek := int(bookingStart.Weekday())

	// Look up availability record for this staff and day
	var availability models.StaffAvailability
	err := db.DB.Where("business_id = ? AND staff_id = ? AND day_of_week = ?", businessID, staffID, dayOfWeek).
		First(&availability).Error

	// If no availability found, staff is not scheduled
	if err != nil {
		return false
	}

	// Parse stored times and create full datetime on booking date
	layout := "15:04"

	// Extract date from booking start time
	y, m, d := bookingStart.Date()

	availStartParsed, err := time.Parse(layout, availability.StartTime)
	if err != nil {
		return false
	}
	// Construct availability start time with booking date
	availStart := time.Date(y, m, d, availStartParsed.Hour(), availStartParsed.Minute(), 0, 0, bookingStart.Location())

	availEndParsed, err := time.Parse(layout, availability.EndTime)
	if err != nil {
		return false
	}
	// Construct availability end time with booking date
	availEnd := time.Date(y, m, d, availEndParsed.Hour(), availEndParsed.Minute(), 0, 0, bookingStart.Location())

	// Check if booking time fits within staff availability
	// bookingStart >= availStart AND bookingEnd <= availEnd
	if (bookingStart.Equal(availStart) || bookingStart.After(availStart)) &&
		(bookingEnd.Equal(availEnd) || bookingEnd.Before(availEnd)) {
		return true
	}

	return false
}

package services

import (
	"errors"
	"time"

	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
)

func SetStaffAvailability(businessID, staffID uint, dayOfWeek int, start, end time.Time) error {
	// Validate start < end
	// Since we care about time of day, we extract HH:mm and parse back to time relative to same day, or simply compare if on same day
	// But start and end might be from arbitrary dates passed by controller.
	// The prompt implies we receive time.Time. Let's compare limits.
	// We'll normalize to just HH:mm comparison logic or assume valid same-day times.
	// Usually controller parses "09:00" to 0000-01-01 09:00:00.

	startStr := start.Format("15:04")
	endStr := end.Format("15:04")

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

func IsStaffAvailableForBooking(businessID, staffID uint, bookingStart, bookingEnd time.Time) bool {
	// Determine day_of_week
	dayOfWeek := int(bookingStart.Weekday())

	// Find availability
	var availability models.StaffAvailability
	err := db.DB.Where("business_id = ? AND staff_id = ? AND day_of_week = ?", businessID, staffID, dayOfWeek).
		First(&availability).Error

	if err != nil {
		return false // No availability found = not working
	}

	// Parse stored times relative to booking date
	layout := "15:04"

	// Create availStart and availEnd on the same date as bookingStart
	// We need year, month, day from bookingStart
	y, m, d := bookingStart.Date()

	availStartParsed, err := time.Parse(layout, availability.StartTime)
	if err != nil {
		return false
	}
	availStart := time.Date(y, m, d, availStartParsed.Hour(), availStartParsed.Minute(), 0, 0, bookingStart.Location())

	availEndParsed, err := time.Parse(layout, availability.EndTime)
	if err != nil {
		return false
	}
	availEnd := time.Date(y, m, d, availEndParsed.Hour(), availEndParsed.Minute(), 0, 0, bookingStart.Location())

	// Check if booking fits
	// bookingStart >= availStart AND bookingEnd <= availEnd
	if (bookingStart.Equal(availStart) || bookingStart.After(availStart)) &&
		(bookingEnd.Equal(availEnd) || bookingEnd.Before(availEnd)) {
		return true
	}

	return false
}

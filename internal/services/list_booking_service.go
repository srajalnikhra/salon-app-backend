package services

import (
	"time"

	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
)

// ListBookings retrieves bookings with optional filters
// Can filter by date and/or staff member
func ListBookings(
	businessID uint,
	date *time.Time,
	staffID *uint,
) ([]models.Booking, error) {

	var bookings []models.Booking

	// Build query with base filters
	query := db.DB.
		Where("business_id = ?", businessID).
		Order("start_time ASC")

	// Filter by date if provided
	if date != nil {
		start := date.Truncate(24 * time.Hour)
		end := start.Add(24 * time.Hour)

		query = query.Where(
			"start_time >= ? AND start_time < ?",
			start,
			end,
		)
	}

	// Filter by staff member if provided
	if staffID != nil {
		query = query.Where("staff_id = ?", *staffID)
	}

	err := query.Find(&bookings).Error
	return bookings, err
}

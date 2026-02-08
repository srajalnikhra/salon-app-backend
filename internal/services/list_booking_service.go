package services

import (
	"time"

	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
)

func ListBookings(
	businessID uint,
	date *time.Time,
	staffID *uint,
) ([]models.Booking, error) {

	var bookings []models.Booking

	query := db.DB.
		Where("business_id = ?", businessID).
		Order("start_time ASC")

	if date != nil {
		start := date.Truncate(24 * time.Hour)
		end := start.Add(24 * time.Hour)

		query = query.Where(
			"start_time >= ? AND start_time < ?",
			start,
			end,
		)
	}

	if staffID != nil {
		query = query.Where("staff_id = ?", *staffID)
	}

	err := query.Find(&bookings).Error
	return bookings, err
}

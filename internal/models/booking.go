package models

import "time"

// Booking represents a salon service appointment
// Tracks which customer booked what service at what time with which staff
type Booking struct {
	ID         uint      `gorm:"primaryKey"`                // Unique booking identifier
	BusinessID uint      `gorm:"index;not null"`            // Foreign key: salon (multi-tenant isolation)
	CustomerID uint      `gorm:"index;not null"`            // Foreign key: customer making booking
	ServiceID  uint      `gorm:"index;not null"`            // Foreign key: service being booked
	StaffID    *uint     `gorm:"index"`                     // Foreign key: assigned staff (nullable, admin assigns later)
	Status     string    `gorm:"type:varchar(20);not null"` // Booking status (pending, confirmed, completed, cancelled)
	StartTime  time.Time `gorm:"index;not null"`            // Appointment start time
	EndTime    time.Time `gorm:"index;not null"`            // Appointment end time (calculated from service duration)
	IsActive   bool      `gorm:"default:true"`              // Booking status flag
	Notes      string    // Additional notes from customer
	CreatedAt  time.Time // Booking creation timestamp
	UpdatedAt  time.Time // Last update timestamp
}

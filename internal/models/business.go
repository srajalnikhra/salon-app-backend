package models

import "time"

// Business represents a salon business (multi-tenant support)
// Each business is owned by one admin and contains staff, services, and bookings
type Business struct {
	ID        uint      `gorm:"primaryKey"`     // Unique business identifier
	AdminID   uint      `gorm:"index;not null"` // Foreign key: owner admin
	Name      string    `gorm:"not null"`       // Salon business name
	Phone     string    // Contact phone number
	Address   string    // Physical location
	Timezone  string    // Business timezone for scheduling
	CreatedAt time.Time // Business creation timestamp
	UpdatedAt time.Time // Last update timestamp
	IsActive  bool      `gorm:"default:true"` // Business active status
}

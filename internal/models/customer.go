package models

import "time"

// Customer represents a client booking salon services
// Customers are created automatically when bookings are made
type Customer struct {
	ID         uint      `gorm:"primaryKey"`     // Unique customer identifier
	BusinessID uint      `gorm:"index;not null"` // Foreign key: salon they book at
	Name       string    // Full name
	Phone      string    `gorm:"index;not null"` // Primary contact number (for lookups)
	Email      string    // Email address
	IsActive   bool      `gorm:"default:true"` // Customer status
	CreatedAt  time.Time // First booking date
	UpdatedAt  time.Time // Last update timestamp
}

package models

import "time"

// Staff represents a salon employee (barber, stylist, etc.)
// Staff members can offer services and have availability schedules
type Staff struct {
	ID         uint      `gorm:"primaryKey"`      // Unique staff identifier
	BusinessID uint      `gorm:"index;not null"`  // Foreign key: salon they work for
	Name       string    `gorm:"not null"`        // Full name
	Phone      string    `gorm:"not null"`        // Contact phone number
	PIN        string    `gorm:"not null"`        // Hashed PIN for login
	Role       string    `gorm:"default:'staff'"` // Position/role (barber, stylist, etc.)
	IsActive   bool      `gorm:"default:true"`    // Employment status
	CreatedAt  time.Time // Hire date/record creation
	UpdatedAt  time.Time // Last update timestamp
}

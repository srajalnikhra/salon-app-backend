package models

import "time"

// Admin represents a salon business administrator account
// Each admin manages one or more salons and can create/manage staff and bookings
type Admin struct {
	ID        uint      `gorm:"primaryKey"`           // Unique admin identifier
	Name      string    `gorm:"not null"`             // Full name of the admin
	Email     string    `gorm:"uniqueIndex;not null"` // Unique email for login
	Password  string    `gorm:"not null"`             // Hashed password (bcrypt)
	IsActive  bool      `gorm:"default:true"`         // Account status
	CreatedAt time.Time // Account creation timestamp
	UpdatedAt time.Time // Last update timestamp
}

package models

import "time"

// StaffService represents a many-to-many relationship between Staff and Services
// Defines which services each staff member is qualified to offer
type StaffService struct {
	ID         uint      `gorm:"primaryKey"`     // Unique relationship identifier
	BusinessID uint      `gorm:"index;not null"` // Foreign key: salon context
	StaffID    uint      `gorm:"index;not null"` // Foreign key: staff member
	ServiceID  uint      `gorm:"index;not null"` // Foreign key: service offered
	CreatedAt  time.Time // Assignment creation timestamp
}

package models

import "time"

// Service represents a salon service offering (haircut, shave, etc.)
// Each service has a duration and price, and can be assigned to multiple staff
type Service struct {
	ID         uint      `gorm:"primaryKey"`     // Unique service identifier
	BusinessID uint      `gorm:"index;not null"` // Foreign key: salon offering service
	Name       string    `gorm:"not null"`       // Service name
	Duration   int       `gorm:"not null"`       // Service duration in minutes
	Price      float64   `gorm:"not null"`       // Service price
	IsActive   bool      `gorm:"default:true"`   // Service availability
	CreatedAt  time.Time // Service creation timestamp
	UpdatedAt  time.Time // Last update timestamp
}

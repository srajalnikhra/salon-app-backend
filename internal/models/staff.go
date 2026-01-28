package models

import "time"

type Staff struct {
	ID         uint      `gorm:"primaryKey"`
	BusinessID uint      `gorm:"index;not null"` // tenant isolation
	Name       string    `gorm:"not null"`
	Phone      string    `gorm:"not null"`
	Role       string    // e.g. barber, stylist, receptionist
	IsActive   bool      `gorm:"default:true"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

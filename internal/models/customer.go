package models

import "time"

type Customer struct {
	ID         uint      `gorm:"primaryKey"`
	BusinessID uint      `gorm:"index;not null"` // tenant isolation
	Name       string
	Phone      string    `gorm:"index;not null"`
	Email      string
	IsActive   bool      `gorm:"default:true"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

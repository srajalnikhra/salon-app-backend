package models

import "time"

type Service struct {
	ID         uint      `gorm:"primaryKey"`
	BusinessID uint      `gorm:"index;not null"` // tenant isolation
	Name       string    `gorm:"not null"`
	Duration   int       `gorm:"not null"` // duration in minutes
	Price      float64   `gorm:"not null"`
	IsActive   bool      `gorm:"default:true"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

package models

import "time"

type Staff struct {
	ID         uint      `gorm:"primaryKey"`
	BusinessID uint      `gorm:"index;not null"`
	Name       string    `gorm:"not null"`
	Phone      string    `gorm:"not null"`
	PIN        string    `gorm:"not null"` // hashed
	IsActive   bool      `gorm:"default:true"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

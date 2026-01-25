package models

import "time"

type Business struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Phone     string    `gorm:"not null"`
	Address   string
	OwnerID   uint      `gorm:"not null"` // Admin ID

	CreatedAt time.Time
	UpdatedAt time.Time
}

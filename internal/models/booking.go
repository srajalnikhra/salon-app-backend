package models

import "time"

type Booking struct {
	ID uint `gorm:"primaryKey"`

	BusinessID uint `gorm:"index;not null"` // tenant isolation

	CustomerID uint `gorm:"index;not null"`
	ServiceID  uint `gorm:"index;not null"`
	StaffID    *uint `gorm:"index"` // nullable (admin assigns later)

	Status string `gorm:"type:varchar(20);not null"` // pending, confirmed, completed, cancelled

	StartTime time.Time `gorm:"index;not null"`
	EndTime   time.Time `gorm:"index;not null"`

	Notes string

	CreatedAt time.Time
	UpdatedAt time.Time
}

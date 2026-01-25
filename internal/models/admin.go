package models

import "time"

type Admin struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`

	Businesses []Business `gorm:"foreignKey:OwnerID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

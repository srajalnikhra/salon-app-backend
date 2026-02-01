package models

import "time"

type Business struct {
	ID        uint      `gorm:"primaryKey"`
	AdminID   uint      `gorm:"index;not null"`
	Name      string    `gorm:"not null"`
	Phone     string
	Address   string
	Timezone  string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsActive   bool      `gorm:"default:true"`
}

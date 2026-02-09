package models

import "time"

type StaffService struct {
	ID         uint `gorm:"primaryKey"`
	BusinessID uint `gorm:"index;not null"`
	StaffID    uint `gorm:"index;not null"`
	ServiceID  uint `gorm:"index;not null"`
	CreatedAt  time.Time
}

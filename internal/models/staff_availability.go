package models

import "time"

type StaffAvailability struct {
	ID         uint   `gorm:"primaryKey"`
	BusinessID uint   `gorm:"index;not null"`
	StaffID    uint   `gorm:"index;not null"`
	DayOfWeek  int    `gorm:"not null"`                 // 0=Sunday, 6=Saturday
	StartTime  string `gorm:"type:varchar(5);not null"` // HH:mm
	EndTime    string `gorm:"type:varchar(5);not null"` // HH:mm
	CreatedAt  time.Time
}

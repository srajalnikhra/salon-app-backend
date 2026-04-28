package models

import "time"

// StaffAvailability defines working hours for each staff member
// Specifies which days and time slots each staff member works
type StaffAvailability struct {
	ID         uint      `gorm:"primaryKey"`               // Unique availability record identifier
	BusinessID uint      `gorm:"index;not null"`           // Foreign key: salon context
	StaffID    uint      `gorm:"index;not null"`           // Foreign key: staff member
	DayOfWeek  int       `gorm:"not null"`                 // Day of week (0=Sunday, 1=Monday, ..., 6=Saturday)
	StartTime  string    `gorm:"type:varchar(5);not null"` // Working hours start (format: HH:mm)
	EndTime    string    `gorm:"type:varchar(5);not null"` // Working hours end (format: HH:mm)
	CreatedAt  time.Time // Availability setup timestamp
}

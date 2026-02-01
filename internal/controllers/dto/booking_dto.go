package dto

import "time"

type CreateBookingRequest struct {
	BusinessID uint `json:"business_id"`

	Customer struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	} `json:"customer"`

	ServiceID uint  `json:"service_id"`
	StaffID   *uint `json:"staff_id"`

	StartTime time.Time `json:"start_time"`
	Notes     string    `json:"notes"`
}

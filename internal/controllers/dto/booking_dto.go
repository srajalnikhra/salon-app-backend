package dto

import "time"

type CreateBookingRequest struct {
	Customer struct {
		Name  string `json:"name" example:"Rahul"`
		Phone string `json:"phone" example:"9876543210"`
	} `json:"customer"`

	ServiceID uint `json:"service_id" example:"1"`

	StaffID *uint `json:"staff_id" example:"2"`

	StartTime time.Time `json:"start_time" example:"2026-03-03T11:00:00Z"`

	Notes string `json:"notes" example:"Haircut booking"`
}
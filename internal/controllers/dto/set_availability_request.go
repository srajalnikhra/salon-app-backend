package dto

// SetAvailabilityRequest is the DTO for staff availability scheduling
// Defines working hours for a specific day of the week
type SetAvailabilityRequest struct {
	DayOfWeek int    `json:"day_of_week" example:"1"`    // Day (0=Sunday, 1=Monday, ..., 6=Saturday)
	StartTime string `json:"start_time" example:"09:00"` // Work start time (HH:mm format)
	EndTime   string `json:"end_time" example:"18:00"`   // Work end time (HH:mm format)
}

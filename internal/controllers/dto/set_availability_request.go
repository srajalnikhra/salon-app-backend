package dto

type SetAvailabilityRequest struct {
	DayOfWeek int    `json:"day_of_week" example:"1"`
	StartTime string `json:"start_time" example:"09:00"`
	EndTime   string `json:"end_time" example:"18:00"`
}
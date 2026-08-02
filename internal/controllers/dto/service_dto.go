package dto

// CreateServiceRequest is the DTO for admin service creation
// Separates HTTP payload from database model
type CreateServiceRequest struct {
	Name string `json:"name" example:"Hair Cut"` // Service name

	Duration int `json:"duration" example:"30"` // Duration in minutes

	Price float64 `json:"price" example:"250"` // Service price
}

// UpdateServiceRequest is the DTO for updating a service.
// All fields are required because this endpoint performs a full update.
type UpdateServiceRequest struct {
	Name string `json:"name" example:"Hair Spa"` // Updated service name

	Duration int `json:"duration" example:"45"` // Updated duration in minutes

	Price float64 `json:"price" example:"600"` // Updated service price
}
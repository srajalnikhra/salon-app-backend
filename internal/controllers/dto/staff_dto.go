package dto

// CreateStaffRequest is the DTO for admin to create a new staff member
// Validates and structures staff creation request data from HTTP
type CreateStaffRequest struct {
	Name  string `json:"name" example:"Staff One"`   // Staff member full name
	Phone string `json:"phone" example:"9999999999"` // Contact number (used for login)
	PIN   string `json:"pin" example:"1234"`         // PIN for staff login (will be hashed)
	Role  string `json:"role" example:"barber"`      // Position/role in salon
}

package dto

type CreateStaffRequest struct {
	Name  string `json:"name" example:"Staff One"`
	Phone string `json:"phone" example:"9999999999"`
	PIN   string `json:"pin" example:"1234"`
	Role  string `json:"role" example:"barber"`
}

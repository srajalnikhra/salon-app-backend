package dto

// AdminSignupRequest is the request body for admin signup.
type AdminSignupRequest struct {
	Name       string `json:"name" example:"Srajal"`
	Email      string `json:"email" example:"admin@gmail.com"`
	Password   string `json:"password" example:"password123"`
	BusinessID uint   `json:"business_id" example:"1"`
}

// AdminLoginRequest is the request body for admin login.
type AdminLoginRequest struct {
	Email    string `json:"email" example:"admin@gmail.com"`
	Password string `json:"password" example:"password123"`
}
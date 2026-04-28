package controllers

// AuthResponse is the standard response for authentication endpoints
// Contains success flag and JWT token on successful login/signup
type AuthResponse struct {
	Success bool   `json:"success"` // Operation status
	Token   string `json:"token"`   // JWT authentication token
}

// ErrorResponse is the standard error response structure
// Used for API error responses with descriptive messages
type ErrorResponse struct {
	Success bool   `json:"success"` // Operation status (false for errors)
	Message string `json:"message"` // Human-readable error message
}

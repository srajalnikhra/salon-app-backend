package controllers

type AuthResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
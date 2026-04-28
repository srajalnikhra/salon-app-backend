package services

import (
	"errors"

	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
	"github.com/srajalnikhra/salon-app-backend/internal/utils"
)

// AdminSignup creates new admin account with password hashing
// Validates email uniqueness and stores admin in database
func AdminSignup(name, email, password string, businessID uint) (*models.Admin, error) {
	// Check if email already exists
	var existing models.Admin
	if err := db.DB.Where("email = ?", email).First(&existing).Error; err == nil {
		return nil, errors.New("email already registered")
	}

	// Hash password for secure storage
	hashed, _ := utils.HashPassword(password)

	// Create admin record
	admin := models.Admin{
		Name:     name,
		Email:    email,
		Password: hashed,
		IsActive: true,
	}

	db.DB.Create(&admin)
	return &admin, nil
}

// AdminLogin validates admin credentials
// Returns admin details if email and password match
func AdminLogin(email, password string) (*models.Admin, error) {
	// Find admin by email
	var admin models.Admin
	if err := db.DB.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Validate password against hashed password
	if err := utils.CheckPassword(admin.Password, password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Check if admin account is active
	if !admin.IsActive {
		return nil, errors.New("admin is inactive")
	}

	return &admin, nil
}

package services

import (
	"errors"

	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
	"github.com/srajalnikhra/salon-app-backend/internal/utils"
)

func AdminSignup(name, email, password string, businessID uint) (*models.Admin, error) {
	var existing models.Admin
	if err := db.DB.Where("email = ?", email).First(&existing).Error; err == nil {
		return nil, errors.New("email already registered")
	}

	hashed, _ := utils.HashPassword(password)

	admin := models.Admin{
		Name:     name,
		Email:    email,
		Password: hashed,
		IsActive: true,
	}

	db.DB.Create(&admin)
	return &admin, nil
}

func AdminLogin(email, password string) (*models.Admin, error) {
	var admin models.Admin
	if err := db.DB.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := utils.CheckPassword(admin.Password, password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !admin.IsActive {
		return nil, errors.New("admin is inactive")
	}

	return &admin, nil
}

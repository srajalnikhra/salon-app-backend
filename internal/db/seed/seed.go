package seed

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
)

func SeedInitialData() {
	seedAdmin()
	seedBusiness()
	seedServices()
}

func seedAdmin() {
	var count int64
	db.DB.Model(&models.Admin{}).Count(&count)
	if count > 0 {
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	admin := models.Admin{
		Name:     "Super Admin",
		Email:    "admin@salon.com",
		Password: string(password),
		IsActive: true,
	}

	db.DB.Create(&admin)
	log.Println("Admin seeded")
}

func seedBusiness() {
	var count int64
	db.DB.Model(&models.Business{}).Count(&count)
	if count > 0 {
		return
	}

	business := models.Business{
		Name:     "Demo Salon",
		Phone:    "9999999999",
		IsActive: true,
	}

	db.DB.Create(&business)
	log.Println("Business seeded")
}

func seedServices() {
	var count int64
	db.DB.Model(&models.Service{}).Count(&count)
	if count > 0 {
		return
	}

	services := []models.Service{
		{
			BusinessID: 1,
			Name:       "Haircut",
			Duration:   30,
			Price:      200,
			IsActive:   true,
		},
		{
			BusinessID: 1,
			Name:       "Shave",
			Duration:   15,
			Price:      100,
			IsActive:   true,
		},
		{
			BusinessID: 1,
			Name:       "Hair Spa",
			Duration:   45,
			Price:      500,
			IsActive:   true,
		},
	}

	db.DB.Create(&services)
	log.Println("Services seeded")
}

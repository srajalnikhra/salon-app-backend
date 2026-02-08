package seed

import (
	"log"

	"github.com/srajalnikhra/salon-app-backend/internal/db"
	"github.com/srajalnikhra/salon-app-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func Run() {
	admin := seedAdmin()
	seedBusiness(admin.ID)
	seedServices()
}

func seedAdmin() models.Admin {
	var count int64
	db.DB.Model(&models.Admin{}).Count(&count)
	if count > 0 {
		var admin models.Admin
		db.DB.First(&admin)
		return admin
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
	return admin
}

func seedBusiness(adminID uint) {
	var count int64
	db.DB.Model(&models.Business{}).Count(&count)
	if count > 0 {
		return
	}

	business := models.Business{
		AdminID:  adminID,
		Name:     "Demo Salon",
		Phone:    "9999999999",
		Address:  "Demo Address",
		Timezone: "Asia/Kolkata",
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
		{BusinessID: 1, Name: "Haircut", Duration: 30, Price: 200, IsActive: true},
		{BusinessID: 1, Name: "Shave", Duration: 15, Price: 100, IsActive: true},
	}

	db.DB.Create(&services)
	log.Println("Services seeded")
}

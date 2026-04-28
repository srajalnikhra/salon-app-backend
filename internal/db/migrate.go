package db

import (
	"log"

	"github.com/srajalnikhra/salon-app-backend/internal/models"
)

// AutoMigrate creates or updates database tables based on model definitions
// Automatically generates SQL to match the struct fields and GORM tags
func AutoMigrate() {
	// Run migrations for all models in dependency order
	// GORM will auto-create tables and add new columns as needed
	err := DB.AutoMigrate(
		&models.Admin{},
		&models.Business{},
		&models.Staff{},
		&models.Service{},
		&models.Customer{},
		&models.Booking{},
		&models.StaffService{},
		&models.StaffAvailability{},
	)

	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	log.Println("Database migrated successfully")
}

package db

import (
	"log"

	"github.com/srajalnikhra/salon-app-backend/internal/models"
)

func AutoMigrate() {
	err := DB.AutoMigrate(
		&models.Admin{},
		&models.Business{},
	)

	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	log.Println("Database migrated successfully")
}

package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/srajalnikhra/salon-app-backend/internal/config"
)

// DB is the global database connection instance
// Used across all services to query and manipulate data
var DB *gorm.DB

// ConnectGorm establishes connection to PostgreSQL database using GORM ORM
// Creates DSN (Data Source Name) from config and initializes global DB instance
func ConnectGorm(cfg config.DBConfig) {
	// Build PostgreSQL connection string with host, port, user, password, and database name
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	// Open database connection using PostgreSQL driver
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Assign database instance to global DB variable for use throughout application
	DB = database
	log.Println("GORM database connected successfully")
}

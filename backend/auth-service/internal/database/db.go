package database

import (
	"auth-service/internal/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	// Read env file
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatalf("Config Database does't found")
	}

	// Database connection
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect DataBase: %v", err)
	}

	DB.AutoMigrate(models.Users{})
}

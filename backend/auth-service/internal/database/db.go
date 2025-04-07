package database

import (
	"auth-service/internal/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Load env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file")
	}

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

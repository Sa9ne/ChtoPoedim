package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Users struct {
	ID       uint
	Name     string
	Email    string
	Password string
}

func ConnectDB() {
	// Load env file
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatalf("Failed to load .env file")
	}

	// Read env file
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		fmt.Println("Config Database does't found")
	}

	// Database connection
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect DataBase: %v", err)
	}

	db.AutoMigrate(Users{})
}

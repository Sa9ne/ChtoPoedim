package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var DB *pgx.Conn

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Database_url not found")
	}

	DB, err = pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Failed connect database: %v", err)
	}
	fmt.Println("Successful database connect")
}

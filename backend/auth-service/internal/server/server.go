package server

import (
	"LetsEat/backend/auth-service/internal/database"
	"LetsEat/backend/auth-service/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	log.Println("Server auth starting...")

	database.ConnectDB()

	s := gin.Default()

	s.POST("/register", handlers.Register)

	if err := s.Run(":8081"); err != nil {
		log.Fatal("Auth server doesn't start...")
	}
}

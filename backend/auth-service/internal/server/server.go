package server

import (
	"ChtoPoedim/backend/auth-service/internal/database"
	"ChtoPoedim/backend/auth-service/internal/handlers"
	"ChtoPoedim/backend/auth-service/internal/middleware"
	"ChtoPoedim/backend/auth-service/internal/models"
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	database.ConnectDB()

	s := gin.Default()
	s.Use(models.CORSConfig())

	s.POST("/register", handlers.Register)
	s.POST("/login", handlers.Login)
	s.GET("/validate", middleware.RequireAuth, handlers.Validate)

	err := s.Run(":8081")
	if err != nil {
		log.Println("Server doesn't start")
	}
}

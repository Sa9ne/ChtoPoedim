package server

import (
	"LetsEat/backend/auth-service/internal/database"
	"LetsEat/backend/auth-service/internal/handlers"
	"LetsEat/backend/auth-service/internal/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	log.Println("Server auth starting...")

	database.ConnectDB()

	s := gin.Default()

	s.POST("/register", handlers.Register)
	s.POST("/login", handlers.Login)

	protected := s.Group("/")
	protected.Use(middlewares.AuthMiddleware())

	protected.GET("/profile", handlers.Profile)

	if err := s.Run(":8081"); err != nil {
		log.Fatal("Auth server doesn't start...")
	}
}

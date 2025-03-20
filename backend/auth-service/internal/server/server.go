package server

import (
	"LetsEat/backend/auth-service/internal/database"
	"LetsEat/backend/auth-service/internal/handlers"
	"LetsEat/backend/auth-service/internal/middlewares"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	log.Println("Server auth starting...")

	database.ConnectDB()

	s := gin.Default()
	s.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // фронт на порту 8080
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"}, // Разрешаем Authorization
		AllowCredentials: true,
	}))

	s.POST("/api/register", handlers.Register)
	s.POST("/api/login", handlers.Login)

	protected := s.Group("/")
	protected.Use(middlewares.AuthMiddleware())

	protected.GET("/profile", handlers.Profile)

	if err := s.Run(":8081"); err != nil {
		log.Fatal("Auth server doesn't start...")
	}
}

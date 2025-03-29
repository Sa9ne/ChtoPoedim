package server

import (
	"LetsEat/backend/auth-service/internal/database"
	"LetsEat/backend/auth-service/internal/handlers"
	"LetsEat/backend/auth-service/internal/middleware"
	"LetsEat/backend/auth-service/internal/models"
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	database.ConnectDB()

	s := gin.Default()
	s.Use(models.CORSConfig())

	s.GET("/", handlers.WelcomeReq)
	s.POST("api/register", handlers.Register)
	s.POST("api/login", handlers.Login)

	auth := s.Group("/api")
	auth.Use(middleware.AuthCheck())
	auth.GET("/profile", handlers.Profile)

	err := s.Run(":8081")
	if err != nil {
		log.Println("Server doesn't start")
	}
}

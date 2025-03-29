package server

import (
	"LetsEat/backend/auth-service/internal/database"
	"LetsEat/backend/auth-service/internal/handlers"
	"LetsEat/backend/auth-service/internal/middleware"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	database.ConnectDB()

	s := gin.Default()
	s.Use(cors.Default())

	s.GET("/", handlers.WelcomeReq)
	s.POST("/register", handlers.Register)
	s.POST("/login", handlers.Login)

	auth := s.Group("/")
	auth.Use(middleware.AuthCheck())
	auth.GET("/profile", handlers.Profile)

	err := s.Run(":8081")
	if err != nil {
		log.Println("Server doesn't start")
	}
}

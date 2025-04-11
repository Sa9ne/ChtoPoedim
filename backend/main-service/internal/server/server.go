package server

import (
	"log"
	"main-service/internal/database"
	"main-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Start() {
	log.Printf("Starting Server...")
	s := gin.Default()

	database.ConnectDB()

	// Работа статической директории frontend
	s.Static("/frontend", "/Users/user/important/ChtoPoedim/frontend")

	// Маршрутизаторы
	s.GET("/", handlers.WelcomeReq)
	s.GET("/search", handlers.SearchRecipes)

	err := s.Run(":8080")
	if err != nil {
		log.Fatal("Failed to create server")
	}
}

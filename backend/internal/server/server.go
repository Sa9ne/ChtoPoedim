package server

import (
	"LetsEat/backend/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	log.Printf("Starting Server...")
	s := gin.Default()

	// Работа статической директории frontend
	s.Static("/frontend", "D:/important/Let's Eat/frontend")
	// Работа статической иконки favicon
	s.StaticFile("favicon.ico", "D:/important/Let's Eat/frontend/icons/favicon.ico")

	// Маршрутизатоор главной страницы
	s.GET("/", handlers.WelcomeReq)

	err := s.Run(":8080")
	if err != nil {
		log.Fatal("Failed to create server")
	}
}

package server

import (
	"LetsEat/backend/internal/database"
	"LetsEat/backend/internal/handlers"
	"LetsEat/backend/internal/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	log.Printf("Starting Server...")
	s := gin.Default()

	database.ConnectDB()

	// Работа статической директории frontend
	s.Static("/frontend", "D:/important/Let's Eat/frontend")

	// Ставим middlewares для проверки языка
	s.Use(middlewares.LangCheck)

	// Маршрутизаторы
	s.GET("/", handlers.WelcomeReq)
	s.GET("/search", handlers.SearchRecipes)
	s.GET("/set-languages")

	err := s.Run(":8080")
	if err != nil {
		log.Fatal("Failed to create server")
	}
}

package server

import (
	"food-service/internal/database"
	"food-service/internal/handlers"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	s := gin.Default()

	database.ConnectDB()

	s.Use(cors.Default())

	s.GET("/DishDay", handlers.DishDay)

	err := s.Run(":8082")
	if err != nil {
		log.Fatal("Failed to create server")
	}
}

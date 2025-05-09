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

	s.Static("/frontend", "/Users/user/important/ChtoPoedim/frontend")

	s.GET("/DishDay", handlers.DishDay)
	s.GET("/DishCatalog", handlers.Catalog)
	s.POST("/SmartSelectFood", handlers.SmartSelect)

	// Избранные блюда пользователя
	FavDish := s.Group("/favorites")

	FavDish.GET("/", handlers.GetFavorites)
	FavDish.POST("/:DishName", handlers.AddFavorites)
	FavDish.DELETE("/:DishName", handlers.DeleteFavorites)

	err := s.Run(":8082")
	if err != nil {
		log.Fatal("Failed to create server")
	}
}

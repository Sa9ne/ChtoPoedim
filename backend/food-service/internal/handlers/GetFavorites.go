package handlers

import (
	"food-service/internal/database"
	"food-service/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFavorites(ctx *gin.Context) {
	// Проверяем пользователя
	userValue, exist := ctx.Get("user")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Приводим пользователя к нужному виду
	User, ok := userValue.(models.Users)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User type mismatch"})
		return
	}

	// Проверяем есть ли блюда в избранном
	var FavoriteDishes []models.Favorites
	if err := database.DB.Where("user_id = ?", User.ID).Find(&FavoriteDishes).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nothing found"})
	}

	// Если нет избранных блюд
	if len(FavoriteDishes) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No favorites found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"favorite": FavoriteDishes})
}

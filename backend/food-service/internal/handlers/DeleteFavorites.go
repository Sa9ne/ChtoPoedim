package handlers

import (
	"food-service/internal/database"
	"food-service/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteFavorites(ctx *gin.Context) {
	// Получаем блюдо из запроса
	DishFav, err := strconv.Atoi(ctx.Param("DishId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dish id don't found"})
		return
	}

	// Проверяем пользователя
	userValue, exist := ctx.Get("user")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Приводим пользователя к нужному типу
	User, ok := userValue.(models.Users)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Type user problem"})
		return
	}

	// Проверяем есть ли в любимом
	var FavoriteDishes models.Favorites
	if err := database.DB.Where("user_id = ? AND dish_id = ?", User.ID, DishFav).First(&FavoriteDishes).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Dish don't found"})
		return
	}

	if err := database.DB.Delete(&FavoriteDishes).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dish was nor deleted"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successful delete"})
}

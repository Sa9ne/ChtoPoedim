package handlers

import (
	"food-service/internal/database"
	"food-service/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddFavorites(ctx *gin.Context) {
	// Смотрим какое блюдо выбрал пользователь
	DishFav, err := strconv.Atoi(ctx.Param("DishId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dish ID"})
		return
	}

	// Проверяем авторизирован ли пользователь
	userValue, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Приводим пользователя к нужному виду
	User, ok := userValue.(models.Users)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User type mismatch"})
		return
	}

	// Находим это блюдо
	var dish models.Dishes
	if err := database.DB.Where("id = ?", DishFav).First(&dish).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Dish not found"})
		return
	}

	// Проверка на тот случай, если уже добавлен
	var existing models.Favorites
	err = database.DB.Where("UserID = ? AND DishID = ?", User.ID, dish.ID).First(&existing).Error
	if err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Already added"})
		return
	}
	if err != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Сохранение
	Favorite := models.Favorites{
		UserID: User.ID,
		DishID: dish.ID,
	}

	if err := database.DB.Create(&Favorite).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save favorite"})
		return
	}

	// Возвращаем успешный ответ
	ctx.JSON(http.StatusOK, gin.H{"message": "Added to favorites"})
}

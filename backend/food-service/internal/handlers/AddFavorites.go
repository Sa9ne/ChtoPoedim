package handlers

import (
	"food-service/internal/database"
	"food-service/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddFavorites(ctx *gin.Context) {
	// Запросили какое блюдо нужно добавить в избранное
	DishName := ctx.Param("DishName")
	if DishName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Имя блюда не указано"})
		return
	}

	// Получаем пользователя из middleware
	UserInterface, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден в контексте"})
		return
	}
	user, ok := UserInterface.(models.Users)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка преобразования пользователя"})
		return
	}

	// Проверки
	var dish models.Dishes
	if err := database.DB.First(&dish, "name = ?", DishName).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Блюдо не найдено"})
		return
	}

	for _, fav := range user.Favorites {
		if fav.ID == dish.ID {
			ctx.JSON(http.StatusConflict, gin.H{"message": "Блюдо уже в избранном"})
			return
		}
	}

	// Добавляем блюдо в избранное
	if err := database.DB.Model(&user).Association("Favorites").Append(&dish); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении в избранное"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Блюдо добавлено в избранное"})
}

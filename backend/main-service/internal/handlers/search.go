package handlers

import (
	"main-service/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchRecipes(ctx *gin.Context) {
	// Получаем параметр из поиска (?search=pizza)
	query := ctx.Query("search")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Пустое поле ввода"})
		return
	}

	var dishes []database.Dishes
	result := database.DB.Where("name ILIKE ?", "%"+query+"%").Find(&dishes)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Ошибка при поиске", "error": result.Error.Error()})
		return
	}

	if len(dishes) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Блюдо не найдено"})
		return
	}

	ctx.JSON(http.StatusOK, dishes)
}

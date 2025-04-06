package handlers

import (
	"ChtoPoedim/backend/main-service/internal/database"
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
	database.DB.Where("name ILIKE ?", "%"+query+"%").Find(&dishes)

	if len(dishes) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Блюдо не найдено"})
		return
	}

	ctx.JSON(http.StatusOK, dishes)
}

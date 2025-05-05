package handlers

import (
	"food-service/internal/database"
	"food-service/internal/models"
	"math/rand/v2"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DishDay(ctx *gin.Context) {
	var Dishes models.Dishes
	// Считаем количество блюд
	var count int64
	if err := database.DB.Model(&models.Dishes{}).Count(&count).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to count dishes"})
		return
	}

	if count == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "dishes doesn't found"})
		return
	}

	// Рандомизация id
	RandDish := rand.IntN(int(count))

	// Ищем в бд рандомные блюда
	if err := database.DB.Offset(RandDish).Limit(1).Find(&Dishes).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get random dish"})
		return
	}

	ctx.JSON(http.StatusOK, Dishes)
}

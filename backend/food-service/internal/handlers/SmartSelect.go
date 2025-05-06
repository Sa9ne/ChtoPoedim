package handlers

import (
	"food-service/internal/database"
	"food-service/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SmartSelect(ctx *gin.Context) {
	var Choices []models.UserChoice

	// Парсим данные с Frontend
	if err := ctx.Bind(&Choices); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Choice info doesn't found"})
		return
	}

	// Место, где будем формировать sql запрос
	sql := database.DB.Model(&models.Dishes{})

	// Достаем информацию и формируем SQL запрос
	for _, choice := range Choices {
		if choice.Stage == "vegan" {
			sql = sql.Where("vegan = ?", choice.Choice)
		} else if choice.Stage == "spicy" {
			sql = sql.Where("spicy = ?", choice.Choice)
		} else if choice.Stage == "hearty" {
			sql = sql.Where("hearty = ?", choice.Choice)
		}
	}

	// Ищем в бд нужные блюда
	var Catalog []models.Dishes

	if err := sql.Find(&Catalog).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Dish doesn't found"})
		return
	}

	ctx.JSON(http.StatusOK, Catalog)
}

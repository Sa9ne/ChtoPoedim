package handlers

import (
	"food-service/internal/database"
	"food-service/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Catalog(ctx *gin.Context) {
	var Dishes []models.Dishes

	if err := database.DB.Find(&Dishes).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to found dish"})
		return
	}

	ctx.JSON(http.StatusOK, Dishes)
}

package handlers

import (
	"LetsEat/backend/auth-service/internal/database"
	"LetsEat/backend/auth-service/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(ctx *gin.Context) {
	var user models.Users

	userID := ctx.GetHeader("user")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Don't have user header"})
		return
	}

	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can't find user"})
		return
	}
}

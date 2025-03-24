package handlers

import (
	"LetsEat/backend/auth-service/internal/database"
	"LetsEat/backend/auth-service/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var user models.Users

	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Struct doesn't work"})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User doesn't create"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "New user created successfully!"})
}

package handlers

import (
	"LetsEat/backend/auth-service/internal/database"
	"LetsEat/backend/auth-service/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var user models.Users

	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "struct user not found"})
		return
	}

	if err := database.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Email or Password input invalid"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Log In was successful!"})
}

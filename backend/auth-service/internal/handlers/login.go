package handlers

import (
	"LetsEat/backend/auth-service/internal/database"
	"LetsEat/backend/auth-service/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	var input models.Users

	// Парсим данные
	if err := ctx.Bind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "struct user not found"})
		return
	}

	// Проверяем почту
	var user models.Users
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Email input invalid"})
		return
	}

	// Проверяем пароль
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password input invalid"})
		return
	}

	// Генерация jwt токена
	token, err := models.GenerateToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при генерации токена"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

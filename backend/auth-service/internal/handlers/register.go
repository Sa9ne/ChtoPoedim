package handlers

import (
	"LetsEat/backend/auth-service/internal/database"
	"LetsEat/backend/auth-service/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var input models.Users

	// Парсим данные
	if err := ctx.Bind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Struct doesn't work"})
		return
	}

	// Проверяем был ли зарегистрирован данный email
	var existingUser models.Users
	if err := database.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "This email was be registered"})
		return
	}

	// Хешируем пароль
	cost := 10
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), cost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Password was not hashed"})
		return
	}

	// Создаем структуру, где уже пароль захеширован
	user := models.Users{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(HashedPassword),
	}

	// Создаем пользователя в БД, со структурой user
	if err := database.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User doesn't create"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"name":  user.Name,
		"email": user.Email,
		"id":    user.ID,
	})
}

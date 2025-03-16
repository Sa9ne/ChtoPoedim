package handlers

import (
	"LetsEat/backend/auth-service/internal/database"
	"LetsEat/backend/auth-service/internal/models"
	"LetsEat/backend/auth-service/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login - обработчик для входа
func Login(ctx *gin.Context) {
	var input models.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.Println("Ошибка парсинга JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}

	// Проверяем, существует ли пользователь с таким email
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		log.Println("Пользователь не найден:", input.Email)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
		return
	}

	// Сравниваем введенный пароль с хешированным паролем в базе данных
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		log.Println("Неверный пароль")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
		return
	}

	// Генерируем JWT
	token, err := utils.GenerateJWT(user)
	if err != nil {
		log.Println("Ошибка генерации JWT:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при генерации токена"})
		return
	}

	// Возвращаем токен
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

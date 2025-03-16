package handlers

import (
	"LetsEat/backend/auth-service/internal/database"
	"LetsEat/backend/auth-service/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var input models.User

	if database.DB == nil {
		log.Println("Ошибка: База данных не подключена.")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}

	// Парсим данные из тела запроса в структуру User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}

	// Проверяем, существует ли уже такой email
	var existingUser models.User
	if err := database.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Этот email уже зарегистрирован"})
		return
	}

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хешировании пароля"})
		return
	}

	// Добавляем нового пользователя
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении пользователя"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Регистрация успешна"})
}

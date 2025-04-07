package handlers

import (
	"auth-service/internal/database"
	"auth-service/internal/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	// Создаем JWT токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Читаем .env файл
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read .env file"})
		return
	}

	// Заносим токен в куки
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{"message": "Log in was successful"})
}

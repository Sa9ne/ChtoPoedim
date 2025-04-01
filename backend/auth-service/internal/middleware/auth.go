package middleware

import (
	"LetsEat/backend/auth-service/internal/database"
	"LetsEat/backend/auth-service/internal/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(ctx *gin.Context) {
	// Получим JWT из cookie
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to find cookies"})
		return
	}

	// Парсим токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse token"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		// Проверим exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Failed exportation"})
			return
		}

		// Найдем такого пользователя
		var user models.Users
		database.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect user ID"})
			return
		}

		// Закрепим этого пользователя
		ctx.Set("user", user)

		// Продолжим
		ctx.Next()

	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Something went wrong"})
		return
	}
}

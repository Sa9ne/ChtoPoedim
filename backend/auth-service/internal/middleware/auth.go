package middleware

import (
	"LetsEat/backend/auth-service/internal/models"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Отсутствует токен"})
			ctx.Abort()
			return
		}

		// Убираем префикс "Bearer" из токена
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Парсим токен
		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return models.GetJwtSecret(), nil
		})
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			ctx.Abort()
			return
		}

		// Присваиваем данные пользователя в контекст запроса
		ctx.Set("user", claims)
		ctx.Next()
	}
}

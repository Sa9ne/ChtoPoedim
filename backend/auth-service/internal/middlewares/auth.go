package middlewares

import (
	"LetsEat/backend/auth-service/internal/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Отсутствует токен"})
			ctx.Abort()
			return
		}

		// Убираем префикс "Bearer " из токена
		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = tokenString[7:]
		}

		// Парсим токен
		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return utils.GetJwtSecret(), nil
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

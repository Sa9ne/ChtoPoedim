package handlers

import (
	"LetsEat/backend/auth-service/internal/database"
	"LetsEat/backend/auth-service/internal/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Profile(ctx *gin.Context) {
	userClaims, _ := ctx.Get("user")
	claims, ok := userClaims.(*jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка при получении данных пользователя из токена"})
		return
	}

	claimsMap := *claims

	var user models.User
	if err := database.DB.First(&user, claimsMap["sub"]).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось найти пользователя"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

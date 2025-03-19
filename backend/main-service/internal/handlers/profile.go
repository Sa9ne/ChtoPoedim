package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProfileHandler(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось получить данные пользователя"})
		return
	}

	email, _ := ctx.Get("email")

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Данные пользователя",
		"user_id": userID,
		"email":   email,
	})
}

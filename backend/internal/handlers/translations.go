package handlers

import (
	"LetsEat/backend/internal/localization"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTranslations(ctx *gin.Context) {
	lang := ctx.DefaultQuery("lang", "en")

	fmt.Printf("Запрос переводов для языка: %s\n", lang)

	translations, err := localization.GetTranslations(lang)
	if err != nil {
		fmt.Printf("Ошибка загрузки переводов: %v\n", err) // Вывод в консоль сервера
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка загрузки переводов: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, translations)
}

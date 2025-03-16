package handlers

import (
	"LetsEat/backend/main-service/internal/localization"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTranslations(ctx *gin.Context) {
	lang := ctx.DefaultQuery("lang", "en")

	translations, err := localization.GetTranslations(lang)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка загрузки переводов: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, translations)
}

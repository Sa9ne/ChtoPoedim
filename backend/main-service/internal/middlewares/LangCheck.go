package middlewares

import "github.com/gin-gonic/gin"

func LangCheck(ctx *gin.Context) {

	lang := ctx.Query("lang")
	if lang == "" {
		lang, _ = ctx.Cookie("lang")
	}

	if lang == "" {
		lang = "en"
	}

	ctx.Set("lang", lang)
	ctx.Next()
}

package handlers

import "github.com/gin-gonic/gin"

func WelcomeReq(ctx *gin.Context) {
	ctx.File("D:/important/ChtoPoedim/frontend/index.html")
}

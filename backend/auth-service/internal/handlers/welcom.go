package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeReq(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello, that just my server")
}

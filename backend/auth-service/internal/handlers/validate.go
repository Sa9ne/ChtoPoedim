package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{"message": user})
}

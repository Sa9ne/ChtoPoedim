package handlers

import (
	"LetsEat/backend/auth-service/internal/database"
	"LetsEat/backend/auth-service/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Profile(ctx *gin.Context) {
	var user models.Users

	userID, _ := ctx.Get("user")
	claims, _ := userID.(*jwt.MapClaims)

	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Don't have user header"})
		return
	}

	claimsMap := *claims

	if err := database.DB.First(&user, claimsMap["sub"]).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can't find user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

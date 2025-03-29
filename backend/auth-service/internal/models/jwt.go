package models

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetJwtSecret() []byte {
	// Получаем секретный ключ из переменной окружения
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET_KEY не задан в .env файле")
	}
	return []byte(jwtSecret)
}

func GenerateToken(user Users) (string, error) {
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Время жизни токена (24 часа)
	}

	//  Создаем новый токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен
	tokenString, err := token.SignedString(GetJwtSecret())
	if err != nil {
		log.Println("Ошибка подписания токена:", err)
		return "", err
	}

	return tokenString, nil
}

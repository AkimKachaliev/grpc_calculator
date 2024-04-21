package utils

import (
	"time"
)

var jwtKey = []byte("secret")

// GenerateJWT генерирует JWT токен на основе ID пользователя.
func GenerateJWT(userID int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

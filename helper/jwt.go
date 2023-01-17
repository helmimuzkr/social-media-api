package helper

import (
	"log"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	strToken, err := token.SigningString()
	if err != nil {
		log.Println("Failed generate token", err)
		return "", err
	}

	return strToken, nil
}

func ExtractToken(t interface{}) int {
	user := t.(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userID"].(float64)
		return int(userId)
	}
	return -1
}

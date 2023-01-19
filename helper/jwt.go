package helper

import (
	"log"
	"social-media-app/config"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	strToken, err := token.SignedString([]byte(config.JWT_KEY))
	if err != nil {
		log.Println("Failed generate token", err)
		return "", err
	}

	return strToken, nil
}

func ExtractToken(t interface{}) int {
	user := t.(*jwt.Token)
	userID := -1
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		switch claims["userID"].(type) {
		case float64:
			userID = int(claims["userID"].(float64))
		case int:
			userID = claims["userID"].(int)
		}
	}

	return userID
}

func ValidateToken(strToken string) *jwt.Token {
	// Decode rawToken, parse from rawToken to jwt.Token
	token, _ := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_KEY), nil
	})

	return token
}

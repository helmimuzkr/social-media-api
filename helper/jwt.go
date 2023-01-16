package helper

import (
	"log"
	"social-media-app/config"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(id uint) (string, error) {
	// create claim
	var claims jwt.MapClaims
	claims["authorized"] = true
	claims["id"] = id

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// generate token to string
	strToken, err := token.SigningString()
	if err != nil {
		log.Println("Failed generate token", err)
		return "", err
	}

	return strToken, nil
}

func ExtractToken(token string) (uint, error) {
	res, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_KEY), nil
	})
	if err != nil {
		log.Println("Failed extract token", err)
		return 0, err
	}

	if res.Valid {
		claim := res.Claims.(jwt.MapClaims)
		id := claim["id"].(float64)

		return uint(id), nil
	}

	return 0, nil
}

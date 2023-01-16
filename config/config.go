package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JWT_KEY string

type AppConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
}

func GetConfig() *AppConfig {
	var appConfig AppConfig

	if err := godotenv.Load("app.env"); err != nil {
		log.Println("Load env failed", err)
		return &AppConfig{}
	}

	appConfig.DBHost = os.Getenv("DB_HOST")
	appConfig.DBPass = os.Getenv("DB_PASS")
	appConfig.DBHost = os.Getenv("DB_HOST")
	appConfig.DBPort = os.Getenv("DB_PORT")
	appConfig.DBName = os.Getenv("DB_NAME")

	JWT_KEY = os.Getenv("JWT_KEY")

	return &appConfig
}

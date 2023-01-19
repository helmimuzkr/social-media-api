package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	JWT_KEY                string
	CloudinaryName         string
	CloudinaryApiKey       string
	CloudinaryApiScret     string
	CloudinaryUploadFolder string
)

type AppConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort int
	DBName string
}

func GetConfig() *AppConfig {
	var appConfig AppConfig

	isRead := true

	if val, found := os.LookupEnv("JWT_KEY"); found {
		JWT_KEY = val
		isRead = false
	}
	if val, found := os.LookupEnv("DB_USER"); found {
		appConfig.DBUser = val
		isRead = false
	}
	if val, found := os.LookupEnv("DB_PASS"); found {
		appConfig.DBPass = val
		isRead = false
	}
	if val, found := os.LookupEnv("DB_HOST"); found {
		appConfig.DBHost = val
		isRead = false
	}
	if val, found := os.LookupEnv("DB_PORT"); found {
		cnv, _ := strconv.Atoi(val)
		appConfig.DBPort = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("DB_NAME"); found {
		appConfig.DBName = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_CLOUD_NAME"); found {
		CloudinaryName = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_API_KEY"); found {
		CloudinaryApiKey = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_API_SECRET"); found {
		CloudinaryApiScret = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_UPLOAD_FOLDER"); found {
		CloudinaryUploadFolder = val
		isRead = false
	}
	if val, found := os.LookupEnv("JWT_KEY"); found {
		JWT_KEY = val
		isRead = false
	}

	if isRead {
		if err := godotenv.Load("app.env"); err != nil {
			log.Println("Load env failed", err)
			return nil
		}

		appConfig.DBUser = os.Getenv("DB_USER")
		appConfig.DBName = os.Getenv("DB_NAME")
		appConfig.DBPass = os.Getenv("DB_PASS")
		appConfig.DBHost = os.Getenv("DB_HOST")
		readData := os.Getenv("DB_PORT")
		appConfig.DBPort, _ = strconv.Atoi(readData)

		JWT_KEY = os.Getenv("JWT_KEY")

		CloudinaryName = os.Getenv("CLOUDINARY_CLOUD_NAME")
		CloudinaryApiKey = os.Getenv("CLOUDINARY_API_KEY")
		CloudinaryApiScret = os.Getenv("CLOUDINARY_API_SECRET")
		CloudinaryUploadFolder = os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
	}

	return &appConfig
}

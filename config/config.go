package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	JWT_KEY string
<<<<<<< HEAD

	CloudinaryName         string
	CloudinaryApiKey       string
	CloudinaryApiScret     string
	CloudinaryUploadFolder string
=======
>>>>>>> ce27822 (feat: perbaikan salah variabel)
)

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
		return nil
	}

	appConfig.DBUser = os.Getenv("DB_USER")
	appConfig.DBName = os.Getenv("DB_NAME")
	appConfig.DBPass = os.Getenv("DB_PASS")
	appConfig.DBHost = os.Getenv("DB_HOST")
	appConfig.DBPort = os.Getenv("DB_PORT")

	JWT_KEY = os.Getenv("JWT_KEY")

	CloudinaryName = os.Getenv("CLOUDINARY_CLOUD_NAME")
	CloudinaryApiKey = os.Getenv("CLOUDINARY_API_KEY")
	CloudinaryApiScret = os.Getenv("CLOUDINARY_API_SECRET")
	CloudinaryUploadFolder = os.Getenv("CLOUDINARY_UPLOAD_FOLDER")

	return &appConfig
}

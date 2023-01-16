package config

import (
	"fmt"
	"log"
	postModel "social-media-app/feature/post/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB(cfg *AppConfig) *gorm.DB {
	var db *gorm.DB

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Open connection gorm failed", err)
		return nil
	}

	return db
}

func GormMigrartion(db *gorm.DB) {
	if err := db.AutoMigrate(postModel.Post{}); err != nil {
		log.Fatal(err)
	}
}

package config

import (
	"fmt"
	"log"
	commentModel "social-media-app/feature/comment/repository"
	postModel "social-media-app/feature/post/repository"
	userModel "social-media-app/feature/user/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB(cfg *AppConfig) *gorm.DB {
	var db *gorm.DB

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection error", err.Error())
		return nil
	}
	return db
}

func GormMigration(db *gorm.DB) {
	if err := db.AutoMigrate(postModel.Post{}); err != nil {
		log.Fatal(err)
		return
	}
	if err := db.AutoMigrate(commentModel.Comment{}); err != nil {
		log.Fatal(err)
		return
	}
	if err := db.AutoMigrate(userModel.User{}); err != nil {
		log.Fatal(err)
		return
	}

}

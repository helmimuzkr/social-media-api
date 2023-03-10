package repository

import (
	comment "social-media-app/feature/comment/repository"
	post "social-media-app/feature/post/repository"
	"social-media-app/feature/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"not_null"`
	LastName  string `gorm:"not_null"`
	Email     string `gorm:"not_null;unique"`
	Password  string `gorm:"not_null"`
	Avatar    string
	Posts     []post.Post       `gorm:"foreignKey:UserID"`
	Comments  []comment.Comment `gorm:"foreignKey:UserID"`
}

// Input API dalam bentuk User (isi database), output API dalam bentuk Core (JSON)
func UserToCore(data User) user.Core {
	return user.Core{
		ID:        data.ID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  data.Password,
		Avatar:    data.Avatar,
	}
}

// Input API dalam bentuk Core (JSON), output API dalam bentuk User (Database)
func CoreToUser(data user.Core) User {
	return User{
		Model:     gorm.Model{ID: data.ID},
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  data.Password,
		Avatar:    data.Avatar,
	}
}

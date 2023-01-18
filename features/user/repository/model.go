package repository

import (
	"social-media-app/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
	Avatar    string
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

// Tempat menentukan error
// Meng-convert pesan error yang sistematis menjadi manusiawi

package services

import (
	"errors"
	"log"
	"social-media-app/features/user"

	"strings"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	qry user.UserRepository
}

func New(ur user.UserRepository) user.UserService {
	return &userService{
		qry: ur,
	}
}

func (us *userService) RegisterServ(newUser user.Core) (user.Core, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("bcrypt error ", err.Error())
		return user.Core{}, errors.New("password process error")
	}
	newUser.Password = string(hashed)

	res, err := us.qry.RegisterRepo(newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicate") { // Kalau error mengandung kata "duplicate"
			msg = "data sudah terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}

// func (us *userService) LoginServ(email, password string) (string, user.Core, error) {
	
// }

// func (us *userService) ProfileServ(token interface{}) (user.Core, error) {
	
// }

// func (us *userService) UpdateServ(token interface{}, updateUser user.Core) (user.Core, error) {
	
// }

// func (us *userService) RemoveServ(token interface{}) error {
	
// }


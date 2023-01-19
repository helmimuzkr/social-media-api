// Tempat menentukan error
// Meng-convert pesan error yang sistematis menjadi manusiawi

package service

import (
	"errors"
	"log"
	"mime/multipart"
	"social-media-app/config"
	"social-media-app/feature/user"
	"social-media-app/helper"

	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	qry user.UserRepository
	vld *validator.Validate
}

func New(ur user.UserRepository, v *validator.Validate) user.UserService {
	return &userService{
		qry: ur,
		vld: v,
	}
}

func (us *userService) RegisterServ(newUser user.Core) (user.Core, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("bcrypt error ", err.Error())
		return user.Core{}, errors.New("password process error")
	}
	newUser.Password = string(hashed)

	err = us.vld.Struct(newUser)
	if err != nil {
		log.Println(err)
		helper.ValidationErrorHandle(err)
	}

	res, err := us.qry.RegisterRepo(newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "Duplicate") {
			msg = "data sudah terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}

func (us *userService) LoginServ(email, password string) (string, user.Core, error) {
	res, err := us.qry.LoginRepo(email)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return "", user.Core{}, errors.New(msg)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password)); err != nil { // res.Password = password di database, password = password input
		log.Println("login compare", err.Error())
		return "", user.Core{}, errors.New("password tidak sesuai")
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = res.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 1 hour
	hashToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := hashToken.SignedString([]byte(config.JWT_KEY))

	return token, res, nil
}

func (us *userService) ProfileServ(token interface{}) (user.Core, error) {
	id := uint(helper.ExtractToken(token))
	log.Println(id)
	if id <= 0 {
		return user.Core{}, errors.New("data tidak ditemukan")
	}

	res, err := us.qry.ProfileRepo(id)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") { // Kalau error mengandung kata "not found"
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}

func (us *userService) GetByIdServ(id uint) (user.Core, error) {
	res, err := us.qry.GetByIdRepo(id)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}

func (us *userService) SearchServ(name string) ([]user.Core, error) {
	res, err := us.qry.SearchRepo(name)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") { // Kalau error mengandung kata "not found"
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return []user.Core{}, errors.New(msg)
	}
	return res, nil
}

func (us *userService) UpdateServ(token interface{}, updateUser user.Core, fileHeader *multipart.FileHeader) (user.Core, error) {
	id := uint(helper.ExtractToken(token))
	if id <= 0 {
		return user.Core{}, errors.New("data tidak ditemukan")
	}

	if fileHeader != nil {
		file, _ := fileHeader.Open()
		uploadURL, err := helper.UploadFile(file, "/user")
		if err != nil {
			log.Println(err)
			return user.Core{}, errors.New("failed to upload image")
		}
		updateUser.Avatar = uploadURL.SecureURL
	}

	res, err := us.qry.UpdateRepo(id, updateUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") { // Kalau error mengandung kata "duplicated"
			msg = "data sudah terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}

func (us *userService) RemoveServ(token interface{}) error {
	id := uint(helper.ExtractToken(token))
	if id <= 0 {
		return errors.New("data tidak ditemukan")
	}

	err := us.qry.RemoveRepo(id)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") { // Kalau error mengandung kata "not found"
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return errors.New(msg)
	}
	return nil
}

var (
	validate = validator.New()
)

func (us *userService) UpdatePassServ(token interface{}, oldPass string, newPass user.Core) (user.Core, error) {
	id := uint(helper.ExtractToken(token))
	if id <= 0 {
		return user.Core{}, errors.New("data tidak ditemukan")
	}

	log.Println(oldPass)
	log.Println(newPass.Password)

	res, err := us.qry.CheckPass(id)
	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(oldPass)); err != nil {
		log.Println("update compare", err.Error())
		return user.Core{}, errors.New("password tidak sesuai")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newPass.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("bcrypt error ", err.Error())
		return user.Core{}, errors.New("password process error")
	}
	newPass.Password = string(hashed)

	res, err = us.qry.UpdatePassRepo(id, newPass)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") { // Kalau error mengandung kata "duplicated"
			msg = "data sudah terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}

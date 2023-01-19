// Tempat menentukan error
// Meng-convert pesan error yang sistematis menjadi manusiawi

package services

import (
	"errors"
	"log"
	"social-media-app/config"
	"social-media-app/features/user"
	"social-media-app/helper"

	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
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
		if strings.Contains(err.Error(), "Email") {
			return user.Core{}, errors.New("format email salah")
		} else if strings.Contains(err.Error(), "FirstName") {
			return user.Core{}, errors.New("format firstname salah")
		} else if strings.Contains(err.Error(), "LastName") {
			return user.Core{}, errors.New("format lastname salah")
		} else if strings.Contains(err.Error(), "Password") {
			return user.Core{}, errors.New("format password salah")
		} else {
			return user.Core{}, errors.New("format inputan salah")
		}
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

func (us *userService) UpdateServ(token interface{}, updateUser user.Core) (user.Core, error) {
	if updateUser.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(updateUser.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("bcrypt error ", err.Error())
			return user.Core{}, errors.New("password process error")
		}
		updateUser.Password = string(hashed)
	}

	id := uint(helper.ExtractToken(token))
	if id <= 0 {
		return user.Core{}, errors.New("data tidak ditemukan")
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

func (us *userService) FileUpload(file user.FileCore) (string, error) {
	//validate
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, err := helper.ImageUploadHelper(file.File)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func (us *userService) UpdatePassServ(token interface{}, updatePass user.Core, newPass string) (user.Core, error) {
	id := uint(helper.ExtractToken(token))
	if id <= 0 {
		return user.Core{}, errors.New("data tidak ditemukan")
	}

	res, err := us.qry.CheckPass(id)
	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(updatePass.Password)); err != nil { // res.Password = password di database, password = password input
		log.Println("login compare", err.Error())
		return user.Core{}, errors.New("password tidak sesuai")
	}

	if newPass != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
		if err != nil {
			log.Println("bcrypt error ", err.Error())
			return user.Core{}, errors.New("password process error")
		}
		updatePass.Password = string(hashed)
	}

	res, err = us.qry.UpdatePassRepo(id, updatePass)
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

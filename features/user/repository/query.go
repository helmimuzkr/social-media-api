package repository

import (
	"errors"
	"log"
	"social-media-app/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserRepository {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) RegisterRepo(newUser user.Core) (user.Core, error) {
	cnv := CoreToUser(newUser)
	err := uq.db.Create(&cnv).Error
	if err != nil {
		log.Println("Query create error", err.Error())
		return user.Core{}, err
	}
	newUser.ID = cnv.ID
	return newUser, nil
}

func (uq *userQuery) LoginRepo(email string) (user.Core, error) {
	res := User{}
	if err := uq.db.Where("email = ?", email).First(&res).Error; err != nil {
		log.Println("Login query error,", err.Error())
		return user.Core{}, errors.New("incorrect email")
	}
	return UserToCore(res), nil
}

func (uq *userQuery) ProfileRepo(id uint) (user.Core, error) {
	res := User{}
	qry := uq.db.First(&res, id)

	if qry.RowsAffected <= 0 {
		log.Println("No data processed")
		return user.Core{}, errors.New("user not found, no data displayed")
	}

	err := qry.Error
	if err != nil {
		log.Println("Query first error", err.Error())
		return user.Core{}, errors.New("data failed to displayed")
	}
	return UserToCore(res), nil
}

func (uq *userQuery) UpdateRepo(id uint, updateUser user.Core) (user.Core, error) {
	
	return user.Core{}, nil
}

func (uq *userQuery) RemoveRepo(id uint) error {

	return nil
}

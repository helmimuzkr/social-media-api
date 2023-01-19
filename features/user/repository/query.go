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

func (uq *userQuery) SearchRepo(name string) ([]user.Core, error) {
	res := []User{}
	tmp := []user.Core{}
	qry := uq.db.Where("first_name LIKE ?", "%"+name+"%").Or("last_name LIKE ?", "%"+name+"%").Find(&res)

	if qry.RowsAffected <= 0 {
		log.Println("No data processed")
		return []user.Core{}, errors.New("user not found, no data displayed")
	}

	err := qry.Error
	if err != nil {
		log.Println("Query first error", err.Error())
		return []user.Core{}, errors.New("data failed to displayed")
	}
	for _, v := range res {
		tmp = append(tmp, UserToCore(v))
	}
	return tmp, nil
}

func (uq *userQuery) GetByIdRepo(id uint) (user.Core, error) {
	res := User{}
	qry := uq.db.Where("id = ?", id).First(&res)

	if qry.RowsAffected <= 0 {
		log.Println("No data processed")
		return user.Core{}, errors.New("user not found, no data displayed")
	}

	err := qry.Error
	if err != nil {
		log.Println("Query find error", err.Error())
		return user.Core{}, errors.New("data failed to displayed")
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
	log.Println(id)
	cnv := CoreToUser(updateUser)
	qry := uq.db.Model(&User{}).Where("id = ?", id).Updates(&cnv)

	if qry.RowsAffected <= 0 {
		log.Println("No data processed")
		return user.Core{}, errors.New("user not found, no data updated")
	}

	err := qry.Error
	if err != nil {
		log.Println("Query find error", err.Error())
		return user.Core{}, err
	}
	log.Println(cnv)
	return UserToCore(cnv), nil
}

func (uq *userQuery) RemoveRepo(id uint) error {
	qry := uq.db.Delete(&User{}, id)

	if qry.RowsAffected <= 0 {
		log.Println("No data processed")
		return errors.New("user not found, no data deleted")
	}

	err := qry.Error
	if err != nil {
		log.Println("Query delete error", err.Error())
		return errors.New("data failed to delete")
	}
	return nil
}

func (uq *userQuery) CheckPass(id uint) (user.Core, error) {
	res := User{}
	if err := uq.db.Where("id = ?", id).First(&res).Error; err != nil {
		log.Println("Login query error,", err.Error())
		return user.Core{}, errors.New("incorrect id")
	}
	return UserToCore(res), nil
}

func (uq *userQuery) UpdatePassRepo(id uint, updateUser user.Core) (user.Core, error) {
	log.Println(id)
	cnv := CoreToUser(updateUser)
	qry := uq.db.Model(&User{}).Where("id = ?", id).Updates(&cnv)

	if qry.RowsAffected <= 0 {
		log.Println("No data processed")
		return user.Core{}, errors.New("user not found, no data updated")
	}

	err := qry.Error
	if err != nil {
		log.Println("Query find error", err.Error())
		return user.Core{}, err
	}
	log.Println(cnv)
	return UserToCore(cnv), nil
}
package user

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	Password  string
	Avatar    string
}

type UserHandler interface {
	RegisterHand() echo.HandlerFunc
// 	LoginHand() echo.HandlerFunc
// 	ProfileHand() echo.HandlerFunc
// 	UpdateHand() echo.HandlerFunc
// 	RemoveHand() echo.HandlerFunc
}

type UserService interface {
	RegisterServ(newUser Core) (Core, error)
// 	LoginServ(email, password string) (string, Core, error)
// 	ProfileServ(token interface{}) (Core, error)
// 	UpdateServ(token interface{}, updateUser Core) (Core, error)
// 	RemoveServ(token interface{}) error
// 	FileUpload(file FileCore) (string, error)
}

type UserRepository interface {
	RegisterRepo(newUser Core) (Core, error)
// 	LoginRepo(email string) (Core, error)
// 	ProfileRepo(id uint) (Core, error)
// 	UpdateRepo(id uint, updateUser Core) (Core, error)
// 	RemoveRepo(id uint) error
}

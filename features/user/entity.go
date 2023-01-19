package user

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required"`
	Password  string `validate:"required"`
	Avatar    string
}

type FileCore struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}

type UserHandler interface {
	RegisterHand() echo.HandlerFunc
	LoginHand() echo.HandlerFunc
	SearchHand() echo.HandlerFunc
	GetByIdHand() echo.HandlerFunc
	ProfileHand() echo.HandlerFunc
	UpdateHand() echo.HandlerFunc
	UpdatePassHand() echo.HandlerFunc
	RemoveHand() echo.HandlerFunc
}

type UserService interface {
	RegisterServ(newUser Core) (Core, error)
	LoginServ(email, password string) (string, Core, error) // Email untuk decrypt
	SearchServ(name string) ([]Core, error)
	GetByIdServ(id uint) (Core, error)
	ProfileServ(token interface{}) (Core, error)
	UpdateServ(token interface{}, updateUser Core) (Core, error)
	UpdatePassServ(token interface{}, oldPass Core, newPass string) (Core, error)
	RemoveServ(token interface{}) error
	FileUpload(file FileCore) (string, error)
}

type UserRepository interface {
	RegisterRepo(newUser Core) (Core, error)
	LoginRepo(email string) (Core, error)
	SearchRepo(name string) ([]Core, error)
	GetByIdRepo(id uint) (Core, error)
	ProfileRepo(id uint) (Core, error)
	UpdateRepo(id uint, updateUser Core) (Core, error)
	UpdatePassRepo(id uint, updatePass Core) (Core, error)
	RemoveRepo(id uint) error
	CheckPass(id uint) (Core, error)
}

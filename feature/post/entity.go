package post

import (
	"mime/multipart"

	"github.com/labstack/echo"
)

type Core struct {
	ID        uint
	Caption   string `validate:"required"`
	Image     string `validate:"omitempty,url"`
	PublicID  string
	Author    string
	Avatar    string
	CreatedAt string
	UpdatedAt string
}

type PostHandler interface {
	Create() echo.HandlerFunc
	MyPost() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetByUserID() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	GetAll() echo.HandlerFunc
}

type PostService interface {
	Create(token interface{}, newPost Core, fileHeader *multipart.FileHeader) error
	MyPost(token interface{}) ([]Core, error)
	Update(token interface{}, postID uint, updatePost Core, fileHeader *multipart.FileHeader) error
	Delete(token interface{}, postID uint) error
	GetByUserID(userID uint) ([]Core, error)
	GetByID(postID uint) (Core, error)
	GetAll() ([]Core, error)
}

type PostRepository interface {
	Create(userID uint, newPost Core) error
	MyPost(postID uint) ([]Core, error)
	Update(userID uint, postID uint, updatePost Core) error
	Delete(userID uint, postID uint) error
	GetByUserID(userID uint) ([]Core, error)
	GetByID(postID uint) (Core, error)
	GetAll() ([]Core, error)
}

package post

import (
	"github.com/labstack/echo"
)

type Core struct {
	ID        uint
	Caption   string
	Image     string
	Author    string
	CreatedAt string
}

type PostHandler interface {
	Create() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	MyPost() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type PostService interface {
	Create(post Core) (Core, error)
	GetAll() ([]Core, error)
	MyPost(token interface{}) ([]Core, error)
	Update(token interface{}, post Core) ([]Core, error)
	Delete(token interface{}, postID int) error
}

type PostRepository interface {
	Create(post Core) (Core, error)
	GetAll() ([]Core, error)
	MyPost(userID int) ([]Core, error)
	Update(userID int, postID int, post Core) ([]Core, error)
	Delete(userID int, postID int) error
}

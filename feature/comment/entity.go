package comment

import "github.com/labstack/echo/v4"

type Core struct {
	ID        uint
	Name      string
	Comment   string
	Avatar    string
	CreatedAt string
	UpdatedAt string
}

type CommentHandler interface {
	Add() echo.HandlerFunc
	GetAll() echo.HandlerFunc
}

type CommentService interface {
	Add(token interface{}, postID uint, comment string) error
	GetAll(postID uint) ([]Core, error)
}

type CommentRepository interface {
	Add(userID uint, postID uint, comment string) error
	GetAll(postID uint) ([]Core, error)
}

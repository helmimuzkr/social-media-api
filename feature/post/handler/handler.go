package handler

import (
	"social-media-app/feature/post"

	"github.com/labstack/echo"
)

type postHandler struct {
	srv post.PostService
}

func NewPostHandler(s post.PostService) post.PostHandler {
	return &postHandler{
		srv: s,
	}
}

func (ph *postHandler) Create() echo.HandlerFunc {
	return func(ctx echo.Context) error {

		return nil
	}
}

func (ph *postHandler) GetAll() echo.HandlerFunc {
	return func(ctx echo.Context) error {

		return nil
	}
}

func (ph *postHandler) MyPost() echo.HandlerFunc {
	return func(ctx echo.Context) error {

		return nil
	}
}

func (ph *postHandler) Update() echo.HandlerFunc {
	return func(ctx echo.Context) error {

		return nil
	}
}

func (ph *postHandler) Delete() echo.HandlerFunc {
	return func(ctx echo.Context) error {

		return nil
	}
}

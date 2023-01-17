package handler

import (
	"net/http"
	"social-media-app/feature/post"
	"social-media-app/helper"

	"github.com/jinzhu/copier"
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
	return func(c echo.Context) error {
		token := c.Get("users")

		postReq := PostRequest{}
		if err := c.Bind(&postReq); err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		file, _ := c.FormFile("image")

		newPost := post.Core{}
		copier.Copy(&newPost, &postReq)

		err := ph.srv.Create(token, newPost, file)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusCreated, "success add new post"))
	}
}

func (ph *postHandler) MyPost() echo.HandlerFunc {
	return func(c echo.Context) error {

		return nil
	}
}

func (ph *postHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		return nil
	}
}

func (ph *postHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		return nil
	}
}

func (ph *postHandler) GetByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {

		return nil
	}
}

func (ph *postHandler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {

		return nil
	}
}

func (ph *postHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		return nil
	}
}

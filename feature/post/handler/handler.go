package handler

import (
	"net/http"
	"social-media-app/feature/post"
	"social-media-app/helper"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
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
		token := c.Get("user")

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
		token := c.Get("user")

		res, err := ph.srv.MyPost(token)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		dataRes := ListUserPostResponse{}
		copier.Copy(&dataRes, &res)

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success get all posts", dataRes))
	}
}

func (ph *postHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		idString := c.Param("post_id")
		postID, _ := strconv.Atoi(idString)

		postReq := PostRequest{}
		if err := c.Bind(&postReq); err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		file, _ := c.FormFile("image")

		updatePost := post.Core{}
		copier.Copy(&updatePost, &postReq)

		err := ph.srv.Update(token, uint(postID), updatePost, file)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusCreated, "success update post"))
	}
}

func (ph *postHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		idString := c.Param("post_id")
		postID, _ := strconv.Atoi(idString)

		err := ph.srv.Delete(token, uint(postID))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusCreated, "success delete post data"))
	}
}

func (ph *postHandler) GetByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idString := c.Param("user_id")
		userID, _ := strconv.Atoi(idString)

		res, err := ph.srv.GetByUserID(uint(userID))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		dataRes := ListPostResponse{}
		copier.Copy(&dataRes, &res)

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success get posts", dataRes))
	}
}

func (ph *postHandler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idString := c.Param("post_id")
		postID, _ := strconv.Atoi(idString)

		res, err := ph.srv.GetByID(uint(postID))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		dataRes := PostResponse{}
		copier.Copy(&dataRes, &res)

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success get post", dataRes))
	}
}

func (ph *postHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ph.srv.GetAll()
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		dataRes := ListPostResponse{}
		copier.Copy(&dataRes, &res)

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success get all posts", dataRes))
	}
}

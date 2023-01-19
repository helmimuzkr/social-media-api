package handler

import (
	"net/http"
	"social-media-app/feature/comment"
	"social-media-app/helper"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type commentHandler struct {
	srv comment.CommentService
}

func NewCommentHandler(s comment.CommentService) comment.CommentHandler {
	return &commentHandler{
		srv: s,
	}
}

func (ch *commentHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		commentReq := commentRequest{}
		if err := c.Bind(&commentReq); err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		if err := ch.srv.Add(token, commentReq.PostID, commentReq.Comment); err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusCreated, "success add new comment"))
	}
}

func (ch *commentHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.QueryParam("post_id")
		postID, _ := strconv.Atoi(idStr)

		res, err := ch.srv.GetAll(uint(postID))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		response := ListCommentResponse{}
		copier.Copy(&response, &res)

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success get all comments", response))
	}
}

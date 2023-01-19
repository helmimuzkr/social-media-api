package service

import (
	"errors"
	"social-media-app/feature/comment"
	"social-media-app/helper"
	"strings"

	"github.com/go-playground/validator/v10"
)

type commentService struct {
	repo     comment.CommentRepository
	validate *validator.Validate
}

func NewCommentService(r comment.CommentRepository, v *validator.Validate) comment.CommentService {
	return &commentService{
		repo:     r,
		validate: v,
	}
}

func (cs *commentService) Add(token interface{}, postID uint, newComment string) error {
	userID := helper.ExtractToken(token)
	if userID < 0 {
		return errors.New("token invalid")
	}

	if err := cs.validate.Struct(comment.Core{Comment: newComment}); err != nil {
		msg := helper.ValidationErrorHandle(err)
		return errors.New(msg)
	}

	if err := cs.repo.Add(uint(userID), postID, newComment); err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "post not found"
		} else {
			msg = "internal server error, failed add comment"
		}

		return errors.New(msg)
	}

	return nil
}

func (cs *commentService) GetAll(postID uint) ([]comment.Core, error) {
	res, err := cs.repo.GetAll(postID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "post not found"
		} else {
			msg = "internal server error, failed get all comment"
		}
		return nil, errors.New(msg)
	}

	return res, nil
}

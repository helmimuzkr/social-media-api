package service

import (
	"errors"
	"social-media-app/feature/comment"
	"social-media-app/helper"
	"social-media-app/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAddCommnet(t *testing.T) {
	v := validator.New()
	repo := mocks.NewCommentRepository(t)
	srv := NewCommentService(repo, v)

	t.Run("Success add comment", func(t *testing.T) {
		repo.On("Add", uint(1), uint(1), "halooo").Return(nil).Once()

		str, _ := helper.GenerateToken(uint(1))
		token := helper.ValidateToken(str)
		err := srv.Add(token, uint(1), "halooo")

		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("token invalid", func(t *testing.T) {
		token := jwt.New(jwt.SigningMethodHS256)
		err := srv.Add(token, uint(1), "halooo")

		assert.NotNil(t, err)
		assert.EqualError(t, err, "token invalid")
	})

	t.Run("Post not found", func(t *testing.T) {
		repo.On("Add", uint(1), uint(1), "halooo").Return(errors.New("not found")).Once()

		str, _ := helper.GenerateToken(uint(1))
		token := helper.ValidateToken(str)
		err := srv.Add(token, uint(1), "halooo")

		assert.NotNil(t, err)
		assert.EqualError(t, err, "post not found")
		repo.AssertExpectations(t)
	})

	t.Run("Validation error", func(t *testing.T) {
		str, _ := helper.GenerateToken(uint(1))
		token := helper.ValidateToken(str)
		err := srv.Add(token, uint(1), "")

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "is required")
	})

	t.Run("Server error", func(t *testing.T) {
		repo.On("Add", uint(1), uint(1), "halooo").Return(errors.New("database error")).Once()

		str, _ := helper.GenerateToken(uint(1))
		token := helper.ValidateToken(str)
		err := srv.Add(token, uint(1), "halooo")

		assert.NotNil(t, err)
		assert.EqualError(t, err, "internal server error")
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	v := validator.New()
	repo := mocks.NewCommentRepository(t)
	srv := NewCommentService(repo, v)

	t.Run("Success get all", func(t *testing.T) {
		resRepo := []comment.Core{
			{
				ID:        5,
				Name:      "john",
				Comment:   "hallo",
				Avatar:    "http://cloudinary.com/avatar.jpg",
				CreatedAt: "2023-01-19T18:22:23.518+07:00",
				UpdatedAt: "2023-01-19T18:22:23.518+07:00",
			},
			{
				ID:        6,
				Name:      "doe",
				Comment:   "nice to meet y'all",
				Avatar:    "",
				CreatedAt: "2023-01-19T18:53:41.272+07:00",
				UpdatedAt: "2023-01-19T18:53:41.272+07:00",
			},
		}
		repo.On("GetAll", uint(1)).Return(resRepo, nil).Once()

		actual, err := srv.GetAll(uint(1))

		assert.Nil(t, err)
		assert.NotNil(t, actual)
		for i := range resRepo {
			assert.Equal(t, resRepo[i].ID, actual[i].ID)
		}
		repo.AssertExpectations(t)
	})

	t.Run("Post not found because id post never exist or 0", func(t *testing.T) {
		actual, err := srv.GetAll(uint(0))

		assert.NotNil(t, err)
		assert.EqualError(t, err, "post not found")
		assert.Nil(t, actual)
	})

	t.Run("Post not found", func(t *testing.T) {
		repo.On("GetAll", uint(1)).Return(nil, errors.New("not found")).Once()

		actual, err := srv.GetAll(uint(1))

		assert.NotNil(t, err)
		assert.EqualError(t, err, "post not found")
		assert.Nil(t, actual)
		repo.AssertExpectations(t)
	})

	t.Run("Server error", func(t *testing.T) {
		repo.On("GetAll", uint(1)).Return(nil, errors.New("database error")).Once()

		actual, err := srv.GetAll(uint(1))

		assert.NotNil(t, err)
		assert.EqualError(t, err, "internal server error, failed get all comment")
		assert.Nil(t, actual)
		repo.AssertExpectations(t)
	})
}

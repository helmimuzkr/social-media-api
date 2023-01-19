package service

import (
	"errors"
	"mime/multipart"
	"social-media-app/feature/post"
	"social-media-app/helper"
	"social-media-app/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAddPost(t *testing.T) {
	v := validator.New()
	repo := mocks.NewPostRepository(t)
	srv := NewPostService(repo, v)

	// t.Run("Add Success", func(t *testing.T) {
	// 	input := post.Core{
	// 		Caption:  "Hallo semuanya!!",
	// 		Image:    "acak.jpg",
	// 		PublicID: "acak",
	// 	}
	// 	repo.On("Create", 1, input).Return(nil).Once()

	// 	// body := bytes.NewBuffer([]byte{12, 11, 32, 31, 98})
	// 	// mw := multipart.NewWriter(body)
	// 	// file, _ := mw.CreateFormFile("image", "image.jpg")

	// 	strToken, _ := helper.GenerateToken(1)
	// 	token := helper.ValidateToken(strToken)
	// 	file := &multipart.FileHeader{
	// 		Filename: "acak.jpg",
	// 		Size:     10,
	// 	}
	// 	err := srv.Create(token, post.Core{Caption: input.Caption}, file)

	// 	assert.Nil(t, err)
	// })

	t.Run("Create token invalid", func(t *testing.T) {
		token := jwt.New(jwt.SigningMethodHS256)
		input := post.Core{
			Caption: "Hallo semuanya!!",
		}
		file := multipart.FileHeader{
			Filename: "test.jpg",
			Size:     10,
		}

		err := srv.Create(token, input, &file)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "token invalid")
	})

	t.Run("Create failed to upload", func(t *testing.T) {
		strToken, _ := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		input := post.Core{
			Caption: "Hallo semuanya!!",
		}
		file := multipart.FileHeader{
			Filename: "test.jpg",
			Size:     10,
		}

		err := srv.Create(token, input, &file)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "failed to upload image")
	})
}

func TestMyPost(t *testing.T) {
	v := validator.New()
	repo := mocks.NewPostRepository(t)
	srv := NewPostService(repo, v)

	t.Run("Success list my post", func(t *testing.T) {
		resRepo := []post.Core{
			{
				ID:        1,
				Caption:   "hallo semua",
				Image:     "https://res.cloudinary.com/dnji8pgyl/image/upload/v1674031689/file/post/zsj8rcaj2p5mlhdacqgd.png",
				CreatedAt: "2023-01-18 15:48:07.441 +0700 WIB",
				UpdatedAt: "2023-01-18 15:48:07.441 +0700 WIB",
			},
			{
				ID:        2,
				Caption:   "Pegel",
				Image:     "https://res.cloudinary.com/dnji8pgyl/image/upload/v1674031689/file/post/zsj8rcaj2p5mlhdacqgd.jpeg",
				CreatedAt: "2023-01-18 15:48:07.441 +0700 WIB",
				UpdatedAt: "2023-01-18 15:48:07.441 +0700 WIB",
			},
		}

		repo.On("MyPost", uint(1)).Return(resRepo, nil).Once()

		strToken, _ := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.MyPost(token)

		assert.Nil(t, err)
		for i := range resRepo {
			assert.Equal(t, resRepo[i].ID, actual[i].ID)
			assert.Equal(t, resRepo[i].Caption, actual[i].Caption)
			assert.Equal(t, resRepo[i].Image, actual[i].Image)
		}
	})

	t.Run("Error invalid token", func(t *testing.T) {
		token := &jwt.Token{}
		actual, err := srv.MyPost(token)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "token invalid")
		assert.Nil(t, actual)

	})

	t.Run("Post not found", func(t *testing.T) {
		repo.On("MyPost", uint(1)).Return(nil, errors.New("not found")).Once()

		strToken, _ := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.MyPost(token)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "post not found")
		assert.Nil(t, actual)
	})

	t.Run("Server error", func(t *testing.T) {
		repo.On("MyPost", uint(1)).Return(nil, errors.New("database error")).Once()

		strToken, _ := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.MyPost(token)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "internal server error")
		assert.Nil(t, actual)
	})

	t.Run("No error but returning 0 data", func(t *testing.T) {
		repo.On("MyPost", uint(1)).Return([]post.Core{}, nil).Once()

		strToken, _ := helper.GenerateToken(1)
		token := helper.ValidateToken(strToken)
		actual, err := srv.MyPost(token)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "post not found")
		assert.Nil(t, actual)
	})
}

func TestUpdate(t *testing.T) {}

// func TestDelete(t *testing.T) {
// 	v := validator.New()
// 	repo := mocks.NewPostRepository(t)
// 	srv := NewPostService(repo, v)

// 	t.Run("Success delete post", func(t *testing.T) {
// 		resRepo := post.Core{
// 			ID:        2,
// 			Caption:   "Pegel",
// 			Image:     "https://res.cloudinary.com/dnji8pgyl/image/upload/v1674031689/file/post/zsj8rcaj2p5mlhdacqgd.jpeg",
// 			PublicID:  "file/post/zsj8rcaj2p5mlhdacqgd",
// 			Author:    "Muzakir",
// 			CreatedAt: "2023-01-18 15:48:07.441 +0700 WIB",
// 			UpdatedAt: "2023-01-18 15:48:07.441 +0700 WIB",
// 		}
// 		repo.On("GetByID", uint(2)).Return(resRepo, nil).Once()

// 		repo.On("Delete", uint(1), uint(2)).Return(nil).Once()

// 		strToken, _ := helper.GenerateToken(1)
// 		token := helper.ValidateToken(strToken)
// 		err := srv.Delete(token, uint(2))

// 		assert.Nil(t, err)
// 	})
// }

func TestGetByUserID(t *testing.T) {
	
}
package service

import (
	"errors"
	"log"
	"mime/multipart"
	"social-media-app/feature/post"
	"social-media-app/helper"
	"strings"

	"github.com/go-playground/validator/v10"
)

type postService struct {
	repo     post.PostRepository
	validate *validator.Validate
}

func NewPostService(r post.PostRepository, v *validator.Validate) post.PostService {
	return &postService{
		repo:     r,
		validate: v,
	}
}

func (ps *postService) Create(token interface{}, newPost post.Core, fileHeader *multipart.FileHeader) error {
	userID := helper.ExtractToken(token)
	if userID < 0 {
		return errors.New("token invalid")
	}

	if fileHeader != nil {
		file, _ := fileHeader.Open()
		uploadURL, err := helper.UploadFile(file, "/post")
		if err != nil {
			log.Println(err)
			return errors.New("failed to upload image")
		}
		newPost.Image = uploadURL.SecureURL
		newPost.PublicID = uploadURL.PublicID
	}

	if err := ps.validate.Struct(newPost); err != nil {
		log.Println(err)
		helper.ValidationErrorHandle(err)
		return err
	}

	if err := ps.repo.Create(uint(userID), newPost); err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}

	return nil
}

func (ps *postService) MyPost(token interface{}) ([]post.Core, error) {
	userID := helper.ExtractToken(token)
	if userID < 0 {
		return nil, errors.New("token invalid")
	}

	res, err := ps.repo.MyPost(uint(userID))
	if err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "post not found"
		} else {
			msg = "internal server error"
		}
		return nil, errors.New(msg)
	}

	if len(res) < 1 {
		return nil, errors.New("post not found")
	}

	return res, nil
}

func (ps *postService) Update(token interface{}, postID uint, updatePost post.Core, fileHeader *multipart.FileHeader) error {
	userID := helper.ExtractToken(token)
	if userID < 0 {
		return errors.New("token invalid")
	}

	// Pengecekan, jika ada file yang diupload atau tidak
	if fileHeader != nil {
		// Mencari post, jika tidak ada maka akan langsung mereturn not found
		res, err := ps.repo.GetByID(postID)
		if err != nil {
			log.Println(err)
			return errors.New("post not found")
		}

		// Upload image ke cloud
		file, _ := fileHeader.Open()
		uploadURL, err := helper.UploadFile(file, "/post")
		if err != nil {
			log.Println(err)
			return errors.New("failed to upload image")
		}
		updatePost.Image = uploadURL.SecureURL
		updatePost.PublicID = uploadURL.PublicID

		// Melakukan validasi ulang
		if err := ps.validate.Struct(updatePost); err != nil {
			log.Println(err)
			helper.ValidationErrorHandle(err)
			return err
		}

		// Hapus image sebelumnya di cloudinary
		if err := helper.DestroyFile(res.PublicID); err != nil {
			log.Println(err)
			return errors.New("failed to upload image")
		}
	}

	if err := ps.repo.Update(uint(userID), postID, updatePost); err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "post not found"
		} else {
			msg = "internal server error"
		}
		return errors.New(msg)
	}

	return nil
}

func (ps *postService) Delete(token interface{}, postID uint) error {
	userID := helper.ExtractToken(token)
	if userID < 0 {
		return errors.New("token invalid")
	}

	res, err := ps.repo.GetByID(postID)
	if err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "post not found"
		} else {
			msg = "internal server error"
		}
		return errors.New(msg)
	}

	if err := ps.repo.Delete(uint(userID), postID); err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "post not found"
		} else {
			msg = "internal server error"
		}
		return errors.New(msg)
	}

	if err := helper.DestroyFile(res.PublicID); err != nil {
		log.Println(err)
		return errors.New("failed to delete image")
	}

	return nil
}

func (ps *postService) GetByUserID(userID uint) ([]post.Core, error) {
	res, err := ps.repo.GetByUserID(uint(userID))
	if err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "user not found"
		} else {
			msg = "internal server error"
		}
		return nil, errors.New(msg)
	}

	return res, nil
}

func (ps *postService) GetByID(postID uint) (post.Core, error) {
	res, err := ps.repo.GetByID(postID)
	if err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "not found") {
			msg = "post not found"
		} else {
			msg = "internal server error"
		}
		return post.Core{}, errors.New(msg)
	}

	return res, nil
}

func (ps *postService) GetAll() ([]post.Core, error) {
	res, err := ps.repo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, errors.New("internal server error")
	}

	return res, nil
}

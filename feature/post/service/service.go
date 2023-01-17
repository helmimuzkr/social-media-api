package service

import (
	"errors"
	"log"
	"mime/multipart"
	"social-media-app/feature/post"
	"social-media-app/helper"

	"github.com/go-playground/validator"
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
		uploadURL, err := helper.Upload(file, "/post")
		if err != nil {
			log.Println(err)
			return errors.New("failed upload image")
		}
		newPost.Image = uploadURL
	}

	err := ps.validate.Struct(newPost)
	if err != nil {
		log.Println(err)
		helper.ValidationErrorHandle(err)
		return err
	}

	err = ps.repo.Create(1, newPost)
	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}

	return nil
}

func (ps *postService) MyPost(token interface{}) ([]post.Core, error) {
	return nil, nil
}

func (ps *postService) Update(token interface{}, updatePost post.Core, file *multipart.FileHeader) error {
	return nil
}

func (ps *postService) Delete(token interface{}, postID uint) error {
	return nil
}

func (ps *postService) GetByUserID(userID uint) ([]post.Core, error) {
	return nil, nil
}

func (ps *postService) GetByID(postID uint) ([]post.Core, error) {
	return nil, nil
}

func (ps *postService) GetAll() ([]post.Core, error) {
	return nil, nil
}

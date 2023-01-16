package service

import (
	"social-media-app/feature/post"

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

func (ps *postService) Create(newPost post.Core) (post.Core, error) {
	return post.Core{}, nil
}

func (ps *postService) GetAll() ([]post.Core, error) {
	return nil, nil
}

func (ps *postService) MyPost(token interface{}) ([]post.Core, error) {
	return nil, nil
}

func (ps *postService) Update(token interface{}, updatePost post.Core) ([]post.Core, error) {
	return nil, nil
}

func (ps *postService) Delete(token interface{}, postID int) error {
	return nil
}

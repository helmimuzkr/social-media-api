package repository

import (
	"social-media-app/feature/post"

	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) post.PostRepository {
	return &postRepository{db: db}
}

func (pr *postRepository) Create(newPost post.Core) (post.Core, error) {

	return post.Core{}, nil
}
func (pr *postRepository) GetAll() ([]post.Core, error) {

	return nil, nil
}
func (pr *postRepository) MyPost(userID int) ([]post.Core, error) {

	return nil, nil
}
func (pr *postRepository) Update(userID int, postID int, updatePost post.Core) ([]post.Core, error) {

	return nil, nil
}
func (pr *postRepository) Delete(userID int, postID int) error {

	return nil
}

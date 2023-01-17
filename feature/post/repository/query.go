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

func (pr *postRepository) Create(userID uint, newPost post.Core) error {
	model := ToModel(newPost)
	tx := pr.db.Where("user_id = ?", userID).Create(&model)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (pr *postRepository) MyPost(userID uint) ([]post.Core, error) {

	return nil, nil
}

func (pr *postRepository) Update(userID uint, postID uint, updatePost post.Core) error {

	return nil
}

func (pr *postRepository) Delete(userID uint, postID uint) error {

	return nil

}

func (pr *postRepository) GetByUserID(userID uint) ([]post.Core, error) {

	return nil, nil
}

func (pr *postRepository) GetByID(postID uint) ([]post.Core, error) {

	return nil, nil
}

func (pr *postRepository) GetAll() ([]post.Core, error) {

	return nil, nil
}

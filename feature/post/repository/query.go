package repository

import (
	"errors"
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
	model.UserID = userID

	tx := pr.db.Where("user_id = ?", model.UserID).Create(&model)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (pr *postRepository) MyPost(userID uint) ([]post.Core, error) {
	posts := []Post{}

	tx := pr.db.Where("user_id = ?", userID).Find(&posts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return ToSliceCore(posts), nil
}

func (pr *postRepository) Update(userID uint, postID uint, updatePost post.Core) error {
	model := ToModel(updatePost)
	model.ID = postID
	model.UserID = userID

	tx := pr.db.Model(&Post{}).Where("id = ? AND user_id = ?", model.ID, model.UserID).Updates(&model)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected < 1 {
		return errors.New("no row updated")
	}

	return nil
}

func (pr *postRepository) Delete(userID uint, postID uint) error {
	tx := pr.db.Unscoped().Where("id = ? AND user_id = ?", postID, userID).Delete(&Post{})
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected < 1 {
		return errors.New("no row deleted")
	}

	return nil
}

func (pr *postRepository) GetByUserID(userID uint) ([]post.Core, error) {
	posts := []UserPost{}

	query := "SELECT posts.id, posts.caption, posts.image, posts.public_id, users.first_name, users.last_name, users.avatar, posts.created_at, posts.updated_at FROM posts JOIN users ON users.id = posts.user_id WHERE posts.deleted_at IS NULL AND posts.user_id  = ?"
	tx := pr.db.Raw(query, userID).Find(&posts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return ToSliceCore(posts), nil
}

func (pr *postRepository) GetByID(postID uint) (post.Core, error) {
	userPost := UserPost{}

	query := "SELECT posts.id, posts.caption, posts.image, posts.public_id, users.first_name, users.last_name, users.avatar, posts.created_at, posts.updated_at FROM posts JOIN users ON users.id = posts.user_id WHERE posts.deleted_at IS NULL AND posts.id = ?"
	tx := pr.db.Raw(query, postID).First(&userPost)
	if tx.Error != nil {
		return post.Core{}, tx.Error
	}

	return ToCore(userPost), nil
}

func (pr *postRepository) GetAll() ([]post.Core, error) {
	posts := []UserPost{}

	query := "SELECT posts.id, posts.caption, posts.image, posts.public_id, users.first_name, users.last_name, users.avatar, posts.created_at, posts.updated_at FROM posts JOIN users ON users.id = posts.user_id WHERE posts.deleted_at IS NULL"
	tx := pr.db.Raw(query).Find(&posts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return ToSliceCore(posts), nil
}

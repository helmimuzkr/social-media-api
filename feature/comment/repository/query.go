package repository

import (
	"social-media-app/feature/comment"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) comment.CommentRepository {
	return &commentRepository{db: db}
}

func (cr *commentRepository) Add(userID uint, postID uint, comment string) error {
	model := ToModel(userID, postID, comment)
	tx := cr.db.Create(&model)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (cr *commentRepository) GetAll(postID uint) ([]comment.Core, error) {
	models := []CommentPost{}

	query := "SELECT comments.id, users.avatar, users.first_name, users.last_name, comments.comment, comments.created_at, comments.updated_at FROM comments JOIN users ON users.id = comments.user_id WHERE comments.post_id = ?"
	tx := cr.db.Raw(query, postID).Find(&models)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return ToSliceCore(models), nil
}

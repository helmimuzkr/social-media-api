package repository

import (
	"social-media-app/feature/comment"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  uint
	PostID  uint
	Comment string
}

type CommentPost struct {
	ID        uint
	Avatar    string
	FirstName string
	LastName  string
	Comment   string
	CreatedAt string
	UpdatedAt string
}

func ToCore(model CommentPost) comment.Core {
	return comment.Core{
		ID:        model.ID,
		Avatar:    model.Avatar,
		Name:      model.FirstName + model.LastName,
		Comment:   model.Comment,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func ToModel(u uint, p uint, c string) Comment {
	return Comment{
		UserID:  u,
		PostID:  p,
		Comment: c,
	}
}

func ToSliceCore(models []CommentPost) []comment.Core {
	cores := []comment.Core{}
	for _, v := range models {
		core := comment.Core{
			ID:        v.ID,
			Avatar:    v.Avatar,
			Name:      v.FirstName + v.LastName,
			Comment:   v.Comment,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		cores = append(cores, core)
	}

	return cores
}

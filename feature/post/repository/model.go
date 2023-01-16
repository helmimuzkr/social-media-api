package repository

import (
	"social-media-app/feature/post"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID  uint
	Caption string
	Image   string
}

type MyPost struct {
	ID        uint
	Caption   string
	Image     string
	FirstName string
	LastName  string
	CreatedAt string
}

// Convert from model to core
func ToCore(model Post) post.Core {
	return post.Core{
		ID:        model.ID,
		Caption:   model.Caption,
		Image:     model.Image,
		CreatedAt: model.CreatedAt.String(),
	}
}

// Convert from core to model
func ToModel(core post.Core) Post {
	return Post{
		Model:   gorm.Model{ID: core.ID},
		Caption: core.Caption,
		Image:   core.Image,
	}
}

// Convert from slice of model to slice of core
func ToSliceCore(models []MyPost) []post.Core {
	listPost := []post.Core{}
	for _, v := range models {
		post := post.Core{
			ID:        v.ID,
			Caption:   v.Caption,
			Image:     v.Image,
			Author:    v.FirstName + v.LastName,
			CreatedAt: v.CreatedAt,
		}

		listPost = append(listPost, post)
	}

	return listPost
}

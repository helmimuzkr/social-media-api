package repository

import (
	comment "social-media-app/feature/comment/repository"
	"social-media-app/feature/post"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Caption  string
	Image    string
	PublicID string
	UserID   uint
	Comments []comment.Comment `gorm:"foreignKey:PostID"`
}

type UserPost struct {
	ID        uint
	Caption   string
	Image     string
	PublicID  string
	FirstName string
	LastName  string
	Avatar    string
	CreatedAt string
	UpdatedAt string
}

// Convert from model to core
func ToCore(model interface{}) post.Core {
	postCore := post.Core{}

	switch v := model.(type) {
	case Post:
		postCore.ID = v.ID
		postCore.Caption = v.Caption
		postCore.Image = v.Image
		postCore.CreatedAt = v.CreatedAt.String()
		postCore.UpdatedAt = v.UpdatedAt.String()

	case UserPost:
		postCore.ID = v.ID
		postCore.Caption = v.Caption
		postCore.Image = v.Image
		postCore.PublicID = v.PublicID
		postCore.Author = v.FirstName + v.LastName
		postCore.Avatar = v.Caption
		postCore.CreatedAt = v.CreatedAt
		postCore.UpdatedAt = v.UpdatedAt
	}

	return postCore
}

// Convert from core to model
func ToModel(core post.Core) Post {
	return Post{
		Model:    gorm.Model{ID: core.ID},
		Caption:  core.Caption,
		Image:    core.Image,
		PublicID: core.PublicID,
	}
}

// Convert from slice of model to slice of core
func ToSliceCore(models interface{}) []post.Core {
	listPost := []post.Core{}

	switch values := models.(type) {
	case []UserPost:
		for _, v := range values {
			post := post.Core{
				ID:        v.ID,
				Caption:   v.Caption,
				Image:     v.Image,
				PublicID:  v.PublicID,
				Author:    v.FirstName + v.LastName,
				Avatar:    v.Avatar,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			}

			listPost = append(listPost, post)
		}
	case []Post:
		for _, v := range values {
			post := post.Core{
				ID:        v.ID,
				Caption:   v.Caption,
				Image:     v.Image,
				PublicID:  v.PublicID,
				CreatedAt: v.CreatedAt.String(),
				UpdatedAt: v.UpdatedAt.String(),
			}

			listPost = append(listPost, post)
		}
	}

	return listPost
}

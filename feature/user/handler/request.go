package handler

import (
	"mime/multipart"
	"social-media-app/feature/user"
)

type File struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}

type Url struct {
	Url string `json:"url,omitempty" validate:"required"`
}

type LoginReq struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterReq struct {
	FirstName string `json:"firstname" form:"firstname"`
	LastName  string `json:"lastname" form:"lastname"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
}

type UpdateReq struct {
	Avatar    string `json:"avatar" form:"avatar"`
	FirstName string `json:"firstname" form:"firstname"`
	LastName  string `json:"lastname" form:"lastname"`
	Email     string `json:"email" form:"email"`
}

type UpdatePass struct {
	OldPassword  string `json:"old_password" form:"old_password"`
	NewPassword  string `json:"new_password" form:"new_password"`
}

type SearchReq struct {
	Name string `json:"name" form:"name"`
}

// Semua input request akan di convert ke Core struct
func ToCore(data interface{}) *user.Core {
	res := user.Core{}

	switch data.(type) {
	case LoginReq:
		cnv := data.(LoginReq)
		res.Email = cnv.Email
		res.Password = cnv.Password
	case RegisterReq:
		cnv := data.(RegisterReq)
		res.FirstName = cnv.FirstName
		res.LastName = cnv.LastName
		res.Email = cnv.Email
		res.Password = cnv.Password
	case UpdateReq:
		cnv := data.(UpdateReq)
		res.FirstName = cnv.FirstName
		res.LastName = cnv.LastName
		res.Email = cnv.Email
		res.Avatar = cnv.Avatar
	case UpdatePass:
		cnv := data.(UpdatePass)
		// res.Password = cnv.OldPassword
		res.Password = cnv.NewPassword
	default:
		return nil
	}
	return &res
}
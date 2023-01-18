package handler

import (
	"social-media-app/features/user"
)

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
	Password  string `json:"password" form:"password"`
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
		res.Password = cnv.Password
		res.Avatar = cnv.Avatar
	default:
		return nil
	}
	return &res
}

package handler

type commentRequest struct {
	PostID  uint   `json:"post_id" form:"post_id"`
	UserID  uint   `json:"user_id" form:"user_id"`
	Comment string `json:"comment" form:"comment"`
}

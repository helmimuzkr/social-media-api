package handler

type PostRequest struct {
	Caption string `json:"caption" form:"caption"`
}

package handler

type CommentResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Comment   string `json:"comment"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListCommentResponse []CommentResponse

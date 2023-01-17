package handler

type UserPostResponse struct {
	ID        uint   `json:"id"`
	Caption   string `json:"caption"`
	Image     string `json:"image"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListUserPostResponse []UserPostResponse

type PostResponse struct {
	ID        uint   `json:"id"`
	Caption   string `json:"caption"`
	Image     string `json:"image"`
	Author    string `json:"author"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListPostResponse []PostResponse

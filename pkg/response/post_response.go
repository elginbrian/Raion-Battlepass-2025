package response

import "time"

type GetAllPostsResponse struct {
	Status string `json:"status"`
	Data   []Post `json:"data"`
}

type SearchPostsResponse struct {
	Status string `json:"status"`
	Data   []Post `json:"data"`
}

type GetPostByIDResponse struct {
	Status string `json:"status"`
	Data   Post   `json:"data"`
}

type CreatePostResponse struct {
	Status string `json:"status"`
	Data   Post   `json:"data"`
}

type UpdatePostResponse struct {
	Status string `json:"status"`
	Data   Post   `json:"data"`
}

type DeletePostResponse struct {
	Status string       `json:"status"`
	Data   RegisterData `json:"data"`
}

type DeletePostData struct {
	Message string `json:"message"`
}

type Post struct {
	ID        string       `json:"id"`
	UserID    string     `json:"user_id"`
	Caption   string    `json:"caption"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
package domain

import "time"

type Post struct {
    ID        string    `json:"id"`
    UserID    string    `json:"user_id"`
    Caption   string    `json:"caption,omitempty"`
    ImageURL  string    `json:"image_url,omitempty"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type PostResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Caption   string `json:"caption,omitempty"`
	ImageURL  string `json:"image_url,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
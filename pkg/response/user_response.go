package response

import "time"

type GetAllUsersResponse struct {
	Status string `json:"status"`
	Data   []User `json:"data"`
}

type GetUserByIDResponse struct {
	Status string `json:"status"`
	Data   User   `json:"data"`
}

type UpdateUserResponse struct {
	Status string `json:"status"`
	Data   User   `json:"data"`
}

type DeleteUserResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
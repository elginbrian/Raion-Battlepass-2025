package response

type LoginResponse struct {
	Status string    `json:"status"`
	Data   LoginData `json:"data"`
}

type LoginData struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	Status string       `json:"status"`
	Data   RegisterData `json:"data"`
}

type RegisterData struct {
	Message string `json:"message"`
}

type GetCurrentUserResponse struct {
	Status string `json:"status"`
	Data   User   `json:"data"`
}

type ChangePasswordResponse struct {
	Status string       `json:"status"`
	Data   RegisterData `json:"data"`
}

type ChangePasswordData struct {
	Message string `json:"message"`
}
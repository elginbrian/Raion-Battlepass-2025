package request

type UpdatePostRequest struct {
	Caption string `json:"caption" validate:"required,min=1"`
}
package domain

import "errors"

var ErrNotFound = errors.New("not found")

type ErrorResponse struct {
	Error string `json:"error"`
}
package errors

import "net/http"

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error∫"`
}

func NewBadRequestError(message string) *Error {
	return &Error{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *Error {
	return &Error{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

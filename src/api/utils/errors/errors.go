package errors

import (
	"net/http"
)

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	status  int    `json:"status"`
	message string `json:"message"`
	error   string `json:"error"`
}

func (e *apiError) Status() int {
	return e.status
}

func (e *apiError) Message() string {
	return e.message
}

func (e *apiError) Error() string {
	return e.error
}

func NewNotFoundError(message string) ApiError {
	return &apiError{
		status:  http.StatusNotFound,
		message: message,
	}
}

func NewBadRequestError(message string) ApiError {
	return &apiError{
		status:  http.StatusBadRequest,
		message: message,
	}
}

func NewInternalServerError(message string) ApiError {
	return &apiError{
		status:  http.StatusInternalServerError,
		message: message,
	}
}

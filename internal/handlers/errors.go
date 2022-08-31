package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

var (
	ErrMethodNotAllowed = &ErrorResponse{HTTPStatusCode: 405, StatusText: "Method not allowed"}
	ErrNotFound         = &ErrorResponse{HTTPStatusCode: 404, StatusText: "Resource not found"}
	ErrBadRequest       = &ErrorResponse{HTTPStatusCode: 400, StatusText: "Bad request"}
)

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Bad request",
		ErrorText:      err.Error(),
	}
}

func ErrServerError(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal server error",
		ErrorText:      err.Error(),
	}
}

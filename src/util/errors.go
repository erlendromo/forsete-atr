package util

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse
//
//	@Summary		ErrorResponse
//	@Description	Json-Response on error
type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

func NewErrorResponse(status int, err error) *ErrorResponse {
	return &ErrorResponse{
		Error:  err.Error(),
		Status: status,
	}
}

func ERROR(w http.ResponseWriter, status int, err error) {
	setHeaders(w, status)
	json.NewEncoder(w).Encode(NewErrorResponse(status, err))
}

package rest

import (
	"net/http"

	"./response"
)

// Error sends error HTTP response.
func Error(w http.ResponseWriter, code int, data interface{}) {
	message := response.Message{
		Error:   &response.ErrorMessage{Code: code, Data: data},
		Success: nil,
	}

	response.JSON(w, code, message)
}

// Success sends successful HTTP response.
func Success(w http.ResponseWriter, code int, data interface{}) {
	message := response.Message{
		Error:   nil,
		Success: &response.SuccessMessage{Code: code, Data: data},
	}

	response.JSON(w, code, message)
}

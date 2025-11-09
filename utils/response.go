package utils

import (
	"encoding/json"
	"net/http"
)

// APIResponse is the standard response structure
type APIResponse struct {
	Message    string `json:"message"`
	Data       any    `json:"data"`
	StatusCode int    `json:"status_code"`
}

// JSONResponse sends a standardized JSON response
func JSONResponse(w http.ResponseWriter, status int, data any, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := APIResponse{
		Message:    message,
		Data:       data,
		StatusCode: status,
	}

	json.NewEncoder(w).Encode(response)
}

// JSONData sends a success response with data
func JSONData(w http.ResponseWriter, data any) {
	JSONResponse(w, http.StatusOK, data, "Request processed successfully")
}

// JSONMessage sends a response with only a message
func JSONMessage(w http.ResponseWriter, message string, status int) {
	JSONResponse(w, status, nil, message)
}

package response

import (
	"encoding/json"
	"net/http"
)

// Responder is an interface for responding to HTTP requests.
type Responder interface {
	Json(w http.ResponseWriter, statusCode int)
}

// Response represents the structure of the response.
type Response struct {
	Code        int         `json:"code"`                   // Http code
	Message     string      `json:"message"`                // Error message or success message
	Total       int64       `json:"total,omitempty"`        // Count of total data
	Data        interface{} `json:"data,omitempty"`         // Object
	ErrorFields interface{} `json:"error_fields,omitempty"` // Validation errors
}

// Json writes a JSON response with the provided data, status code, and message.
func (r *Response) Json(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(r); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
	return
}

// New creates a new Response instance with the provided data.
func New(
	httpCode int,
	message string,
	data interface{},
	total int64,
	errorFields interface{},
) Responder {
	return &Response{
		Code:        httpCode,
		Message:     message,
		Total:       total,
		Data:        data,
		ErrorFields: errorFields,
	}
}

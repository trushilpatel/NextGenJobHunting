package exception

import (
	"net/http"
	"time"
)

// CommonException represents a custom error when a resource is not found.
// swagger:model CommonException
type CommonException struct {
	// The time the error occurred
	// @example "2023-10-09T14:48:00Z"
	Timestamp time.Time `json:"timestamp,omitempty"`

	// Error message
	// @example "Resource not found"
	Message string `json:"message,omitempty"`

	// Optional details about the error
	// @example "User ID does not exist"
	ErrorDetails string `json:"error_details,omitempty"`

	// HTTP status code
	// @example 404
	Status int `json:"status,omitempty"`
}

// NewCommonException creates a new NotFoundException with a given message.
func NewCommonException(message string, errorDetails string) *CommonException {
	return &CommonException{
		Timestamp:    time.Now(),
		Message:      message,
		ErrorDetails: errorDetails,
		Status:       http.StatusNotFound,
	}
}

// Error implements the error interface for NotFoundException.
func (e *CommonException) Error() string {
	return e.Message
}

// StatusCode returns the HTTP status code for the NotFoundException.
func (e *CommonException) StatusCode() int {
	return e.Status
}

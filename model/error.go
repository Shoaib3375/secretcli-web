package model

type ErrorResponse struct {
	Code    int    `json:"code"`    // Error code
	Message string `json:"message"` // Detailed error message
}

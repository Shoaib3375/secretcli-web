package model

type ErrorResponse struct {
	Code    int    `json:"code"`    // Error code
	Message string `json:"message"` // Detailed error message
}

type SuccessResponse struct {
	Code    int         `json:"code"`    // Status code
	Message string      `json:"message"` // Success message
	Data    interface{} `json:"data"`    // Optional data payload
}

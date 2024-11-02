package model

type SuccessResponse struct {
	Code    int         `json:"code"`    // Status code
	Message string      `json:"message"` // Success message
	Data    interface{} `json:"data"`    // Optional data payload
}

package model

// @Description Auth registration request payload
type SwaggerAuthRequest struct {
	Name     string `json:"name" example:"John Doe"`
	Email    string `json:"email" example:"john@example.com"`
	Password string `json:"password" example:"secretpass123"`
}

type SwaggerUserLoginRequest struct {
	Email    string `json:"email" example:"john@example.com"`
	Password string `json:"password" example:"secretpass123"`
}

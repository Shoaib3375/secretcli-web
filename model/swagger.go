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

type SwaggerSecretRequest struct {
	Title    string `json:"title" example:"Facebook Credentials"`
	Username string `json:"username" example:"mokhlesur.mahin"`
	Password string `json:"password" example:"mahin"`
	Note     string `json:"note" example:"This is my Facebook Password. Login korben na."`
	Email    string `json:"email" example:"john@example.com"`
	Website  string `json:"website" example:"https://www.facebook.com/"`
}

type SwaggerGeneratePasswordRequest struct {
	Length               int  `json:"length" example:"8"`
	IncludeSpecialSymbol bool `json:"include_special_symbol" example:"true"`
}

type SwaggerSecretDetailRequest struct {
	SecretID int `json:"secret_id" example:"1"`
}

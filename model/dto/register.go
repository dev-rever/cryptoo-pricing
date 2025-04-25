package model

type RegisterRequest struct {
	Account  string `json:"account" binding:"required,min=6,max=16"`
	Password string `json:"password" binding:"required,min=6,max=16"`
	Email    string `json:"email" binding:"required,emailfmt"`
}

type RegisterResponse struct {
	Account string `json:"account" binding:"required"`
	Email   string `json:"email" binding:"required,emailfmt"`
	Token   string `json:"token" binding:"required"`
}

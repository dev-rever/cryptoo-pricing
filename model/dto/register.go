package model

type RegisterRequest struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=16"`
	Email    string `json:"email" binding:"required"`
}

type RegisterResponse struct {
	Account string `json:"account" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Token   string `json:"token" binding:"required"`
}

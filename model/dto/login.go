package model

type LoginRequest struct {
	Account  string `json:"account" binding:"required,min=6,max=16"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}

type LoginResponse struct {
	Token string `json:"token" binding:"required"`
}

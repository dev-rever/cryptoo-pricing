package model

type UserInfo struct {
	ID       int    `json:"id" binding:"required"`
	Account  string `json:"account" binding:"required"`
	Password string `json:"-"`
	Email    string `json:"email" binding:"required"`
}

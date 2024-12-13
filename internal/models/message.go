package models

type Message struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"required,email"`
}

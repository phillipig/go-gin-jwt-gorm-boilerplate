package models

type User struct {
	Id       uint   `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Telefone string `json:"telefone"`
	Login    string `json:"login"`
	Senha    string `json:"senha"`
}

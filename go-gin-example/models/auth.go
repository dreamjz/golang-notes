package models

type Auth struct {
	Model
	Username string `json:"username"'`
	Password string `json:"password"`
}

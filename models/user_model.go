package models

type UserModel struct {
	ModelImpl
	Username string     `json:"username"`
	Password string     `json:"password"`
	Token    TokenModel `json:"token"`
}

package models

type UserModel struct {
	ModelImpl
	Username string     `json:"username"`
	Password string     `json:"password"`
	Token    TokenModel `json:"token"`
}

func (model *UserModel) Key() string {
	return model.Username
}

func (model *UserModel) SetKey(key string) {
	model.Username = key
}

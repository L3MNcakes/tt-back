package models

type UserModel struct {
	ModelImpl
	Username string      `json:"username"`
	Password string      `json:"password"`
	Token    *TokenModel `json:"token"`
}

func (model *UserModel) Key() string {
	return model.Username
}

func (model *UserModel) SetKey(key string) {
	model.Username = key
}

func (model *UserModel) Bucket() string {
	return "users"
}

func (model *UserModel) SecondaryIndexes() map[string]string {
	return map[string]string{
		"token_bin": model.Token.AccessToken,
	}
}

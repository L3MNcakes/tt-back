package models

type TokenModel struct {
	AccessToken string `json:"accessToken"`
	Expires     string `json:"expires"`
}

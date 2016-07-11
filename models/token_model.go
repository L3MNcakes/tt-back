package models

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
)

const TOKEN_BYTES = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const TOKEN_LENGTH = 32

type TokenModel struct {
	AccessToken string `json:"accessToken"`
	Expires     string `json:"expires"`
}

// More research needed on generating cryptographically secure tokens,
// But this should work fine for the time being.
func GenerateToken() TokenModel {
	token := TokenModel{}
	accessToken := make([]byte, TOKEN_LENGTH)

	// Grab TOKEN_LENGTH random bytes from TOKEN_BYTES
	for i := range accessToken {
		accessToken[i] = TOKEN_BYTES[rand.Intn(len(TOKEN_BYTES))]
	}

	// Hash the random string with SHA-256
	hash := sha256.New()
	hash.Write(accessToken)
	hashSum := hash.Sum(nil)

	token.AccessToken = hex.EncodeToString(hashSum)

	return token
}

package models

import "encoding/json"

type Model interface {
	Key() string
	SetKey(string)
	Json() ([]byte, error)
}

type ModelImpl struct {
	key string
}

func (m *ModelImpl) Key() string {
	return m.key
}

func (m *ModelImpl) SetKey(key string) {
	m.key = key
}

func (m *ModelImpl) Json() ([]byte, error) {
	return json.Marshal(m)
}

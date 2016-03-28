package models

type Model interface {
	Key() string
	SetKey(string)
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

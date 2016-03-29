package models

type Model interface {
	Key() string
	SetKey(string)
	Bucket() string
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

func (m *ModelImpl) Bucket() string {
	return ""
}

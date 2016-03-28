package repositories

import (
	"app/models"
)

type UserRepository struct {
	RepositoryImpl
}

func (repo *UserRepository) Bucket() string {
	return "users"
}

func (repo *UserRepository) Model() models.Model {
	return &models.UserModel{}
}

func (repo *UserRepository) Find(key string) (models.Model, error) {
	return FetchModel(key, repo)
}

func (repo *UserRepository) Save(model models.Model) error {
	return SaveModel(model, repo)
}

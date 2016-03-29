package repositories

import (
	"app/models"
	//"github.com/l3mncakes/tt-back/models"
)

type Repository interface {
	Model() models.Model
	SetModel(models.Model)
	Find(string) (models.Model, error)
	Save() error
}

type RepositoryImpl struct {
	model models.Model
}

func (repo *RepositoryImpl) SetModel(model models.Model) {
	repo.model = model
}

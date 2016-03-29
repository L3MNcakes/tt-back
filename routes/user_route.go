package routes

import (
	"app/models"
	"app/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserRoute struct {
	RouterImpl
}

func (route *UserRoute) Path() string {
	return "/user"
}

func (route *UserRoute) HandleGet(w http.ResponseWriter, r *http.Request) {
	repo := repositories.RiakRepositoryImpl{}
	base_model := &models.UserModel{}

	repo.SetModel(base_model)

	if err := repo.FindBySecondaryIndex("token_bin", "1234567890"); err != nil {
		log.Print(err)
	}

	jval, err := json.Marshal(repo.Model())

	if err != nil {
		log.Print(err)
	}

	fmt.Fprintf(w, "%s", jval)
}

func (route *UserRoute) HandlePost(w http.ResponseWriter, r *http.Request) {
	req := &PostUserRequest{}

	if err := DecodeRequest(req, r.Body); err != nil {
		log.Print(err)
	}

	repo := repositories.RiakRepositoryImpl{}

	model := &models.UserModel{}
	model.Username = req.Username
	model.Password = req.Password

	token := &models.TokenModel{}
	token.AccessToken = "1234567890"

	model.Token = token

	repo.SetModel(model)

	if err := repo.Save(); err != nil {
		log.Print(err)
	}
}

type PostUserRequest struct {
	RequestImpl
	Username string `json:username`
	Password string `json:password`
}

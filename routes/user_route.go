package routes

import (
	"app/repositories"
	"log"
	"net/http"
)

type UserRoute struct {
	RouterImpl
}

func (route *UserRoute) Path() string {
	return "/users"
}

func (route *UserRoute) HandlePost(w http.ResponseWriter, r *http.Request) {
	req := &PostUserRequest{}

	if err := DecodeRequest(req, r.Body); err != nil {
		log.Print(err)
	}

	repo := repositories.UserRepository{}
	model := repo.Model()

	log.Println(model)

	if err := repo.Save(model); err != nil {
		log.Print(err)
	}
}

type PostUserRequest struct {
	RequestImpl
	Username string `json:username`
	Password string `json:password`
}

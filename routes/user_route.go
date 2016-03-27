package routes

import (
	"app/models"
	"encoding/json"
	"fmt"
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

	model := &models.UserModel{
		Username: req.Username,
		Password: req.Password,
	}

	response, err := json.Marshal(model)
	if err != nil {
		log.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", response)
}

type PostUserRequest struct {
	RequestImpl
	Username string `json:username`
	Password string `json:password`
}

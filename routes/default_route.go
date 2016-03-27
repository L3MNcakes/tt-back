package routes

import (
	"io"
	"net/http"
)

type DefaultRoute struct {
	RouterImpl
}

func (route *DefaultRoute) Path() string {
	return "/"
}

func (route *DefaultRoute) HandleGet(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Default Route")
}

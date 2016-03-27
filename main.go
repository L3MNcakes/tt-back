package main

import (
	"app/router"
	"io"
	"net/http"
)

var app = &App{
	Routes: []router.Router{
		&DefaultRoute{},
	},
}

func main() {
	router.Routes(app.Routes)

	http.ListenAndServe(LISTEN_ADDR, nil)
}

type DefaultRoute struct {
	router.RouterImpl
}

func (route *DefaultRoute) Path() string {
	return "/"
}

func (route *DefaultRoute) HandleGet(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

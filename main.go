package main

import (
	"io"
	"net/http"
)

var app = &App{
	Routes: []Router{
		&DefaultRoute{},
	},
}

func main() {
	Routes(app.Routes)

	http.ListenAndServe(LISTEN_HOST+":"+LISTEN_PORT, nil)
}

type DefaultRoute struct {
	RouterImpl
}

func (route *DefaultRoute) Path() string {
	return "/"
}

func (route *DefaultRoute) HandleGet(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

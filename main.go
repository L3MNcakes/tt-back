package main

import (
	"app/routes"
	"net/http"
)

var app = &App{
	Routes: []routes.Router{
		&routes.DefaultRoute{},
	},
}

func main() {
	routes.Routes(app.Routes)

	http.ListenAndServe(":"+LISTEN_PORT, nil)
}

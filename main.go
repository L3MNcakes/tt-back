package main

import (
	"net/http"
)

var app = &App{
	Routes: []Router{
		&DefaultRoute{},
	},
}

func main() {
	Routes(app.Routes)

	http.ListenAndServe(":"+LISTEN_PORT, nil)
}

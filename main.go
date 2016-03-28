package main

import (
	"app/config"
	"app/routes"
	"net/http"
)

func main() {
	app_routes := []routes.Router{
		&routes.DefaultRoute{},
		&routes.UserRoute{},
	}

	routes.Routes(app_routes)

	http.ListenAndServe(":"+config.LISTEN_PORT, nil)
}

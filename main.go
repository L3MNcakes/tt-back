package main

import (
	"app/config"
	"app/routes"
	"net/http"
)

func main() {
	// Initialize routes
	app_routes := []routes.Router{
		&routes.DefaultRoute{},
		&routes.UserRoute{},
	}

	routes.Routes(app_routes)

	// Get served
	http.ListenAndServe(":"+config.LISTEN_PORT, nil)
}

package router

import (
	"log"
	"net/http"
)

type Router interface {
	Path() string
	HandleGet(http.ResponseWriter, *http.Request)
	HandlePost(http.ResponseWriter, *http.Request)
	HandlePut(http.ResponseWriter, *http.Request)
	HandleDelete(http.ResponseWriter, *http.Request)
	HandlePatch(http.ResponseWriter, *http.Request)
}

type RouterImpl struct{}

func (router *RouterImpl) HandleGet(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "405: Method Not Allowed", 405)
}

func (route *RouterImpl) HandlePost(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "405: Method Not Allowed", 405)
}

func (route *RouterImpl) HandlePut(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "405: Method Not Allowed", 405)
}

func (route *RouterImpl) HandleDelete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "405: Method Not Allowed", 405)
}

func (route *RouterImpl) HandlePatch(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "405: Method Not Allowed", 405)
}

func Routes(routes []Router) {
	for _, r := range routes {
		if r.Path() == "" {
			log.Println("Router missing a Path field")
			continue
		}

		http.HandleFunc(r.Path(), route(r))
	}
}

func route(rout Router) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			rout.HandleGet(w, r)
		case "POST":
			rout.HandlePost(w, r)
		case "PUT":
			rout.HandlePut(w, r)
		case "DELETE":
			rout.HandleDelete(w, r)
		case "PATCH":
			rout.HandlePatch(w, r)
		default:
			http.NotFound(w, r)
		}
	}
}

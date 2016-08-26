package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

//NewRouter Create new Router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"UserAdd",
		"POST",
		"/user/add",
		UserAdd,
	},
	Route{
		"UserInfo",
		"GET",
		"/user/{userid}",
		UserInfo,
	},
	Route{
		"GoToken",
		"GET",
		"/token",
		GoToken,
	},
}

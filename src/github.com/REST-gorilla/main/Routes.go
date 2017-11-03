package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

//func NewRouter() *mux.Router {
//
//	router := mux.NewRouter().StrictSlash(true)
//	for _, route := range routes {
//		var handler http.Handler
//		handler = route.HandlerFunc
//		handler = Logger(handler, route.Name)
//
//		router.
//		Methods(route.Method).
//			Path(route.Pattern).
//			Name(route.Name).
//			Handler(route.HandlerFunc)
//	}
//
//	return router
//}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
}
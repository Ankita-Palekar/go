package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		if (route.Name != "GetToken") {
			router.
			Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(jwtMiddleware.Handler(route.HandlerFunc))
		} else {
			router.
			Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(route.HandlerFunc)
		}
	}

	return router
}

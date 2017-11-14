package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/Todos", TodoIndex)
	router.POST("/Todos/:todoId", TodoShow)
	router.GET("/getToken", getToken)
	return router
}

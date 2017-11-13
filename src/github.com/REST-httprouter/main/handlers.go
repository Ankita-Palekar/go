package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var mySigningKey = []byte("secret")

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request, parameters httprouter.Params) {
	todoId := parameters.ByName("todoId")
	fmt.Fprintln(w, "Todo show:", todoId)
}

func getToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["name"] = "Ankita Palekar"
	claims["exp"] = time.Now().Add(time.Second * 120).Unix()
	tokenString, _ := token.SignedString(mySigningKey)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(tokenString))
	fmt.Println(claims["exp"] )
	create(tokenString, "HMAC", "1234567")
}

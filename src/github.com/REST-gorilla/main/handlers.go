package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/dgrijalva/jwt-go"
	"github.com/auth0/go-jwt-middleware"
	"time"
)

var mySigningKey = []byte("secret")

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func TodoIndex(w http.ResponseWriter, r *http.Request) {
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

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func GetToken(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["name"] = "Ankita Palekar"
	claims["exp"] = time.Now().Add(time.Second * 3600).Unix()
	tokenString, _ := token.SignedString(mySigningKey)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(tokenString))
	fmt.Println(claims["exp"])
}

var members = []Member{Member{"someuser", "someuser@somedomain.com"}}

type Member struct {
	Login string
	Email string
}

func PostData(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//
	//var m Member
	//b, _ := ioutil.ReadAll(r.Body)
	//json.Unmarshal(b, &m)
	//fmt.Println(r.Body)
	//
	//
	//members = append(members, m)
	//
	//j, _ := json.Marshal(m)
	//w.Write(j)

	decoder := json.NewDecoder(r.Body)
	fmt.Println(decoder)

}

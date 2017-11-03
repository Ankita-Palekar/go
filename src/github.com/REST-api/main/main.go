package main


import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"encoding/json"
)


func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	//====================== 	using httpRouter	=================

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/users/:id/:name", Hello)

	log.Fatal(http.ListenAndServe(":9090", router))


	// ====================== 	using Gorilla ==============
	router1 := mux.NewRouter().StrictSlash(true)
	router1.HandleFunc("/", Index1)
	router1.HandleFunc("/todos", TodoIndex)
	router1.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":9090", router1))

}


func Index1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome Index1!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
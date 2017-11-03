package main

import (
	"fmt"
	"net/http"
	//"strings"
	//"log"
	"log"
)

type MyMux struct {
}

//func sayhelloName(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()  // parse arguments, you have to call this by yourself
//
//	fmt.Println("-------------")
//	fmt.Println(r)
//	fmt.Println("-------------")
//
//
//	fmt.Println(r.Form)  // print form information in server side
//	fmt.Println("path", r.URL.Path)
//	fmt.Println("scheme", r.URL.Scheme)
//	fmt.Println(r.Form["url_long"])
//	for k, v := range r.Form {
//		fmt.Println("key:", k)
//		fmt.Println("val:", strings.Join(v, ""))
//	}
//	fmt.Fprintf(w,  "Hello Ankita!") // send data to client side
//	fmt.Println(r)
//}


func usersData(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Fprintf(w, "Users data exists here")
}

func main(){
	http.HandleFunc("/", sayhelloName) // set router
	http.HandleFunc("/usersData", usersData)
	http.HandleFunc("/usersdata", usersData)
	err := http.ListenAndServe(":9090", nil) // set listen port

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}



func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

//func main() {
//	mux := &MyMux{}
//	http.ListenAndServe(":9090", mux)
//}
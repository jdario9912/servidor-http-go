package main

import (
	"net/http"
)

func main (){
	http.HandleFunc("/", homeController)
	http.HandleFunc("/api", apiController)
	http.ListenAndServe(":3000", nil)
}

func homeController (res http.ResponseWriter, req *http.Request){
	res.Write([]byte("Hola mundo!"))
}

func apiController (res http.ResponseWriter, req *http.Request){
	res.Write([]byte("Api con Golang!"))
}
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	log.Println("Setting up router...")
	router := mux.NewRouter()
	ConfigureRouter(router)
	http.ListenAndServe(":8080", router)
}

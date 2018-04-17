package main

import (
	"log"
	"net/http"

	"github.com/borislav-rangelov/gonet-example/handlers/home"
	"github.com/borislav-rangelov/gonet-example/handlers/users"
	"github.com/borislav-rangelov/gonet/handlers"
	"github.com/gorilla/mux"
)

func main() {

	log.Println("Setting up router...")
	router := mux.NewRouter()
	configureRouter(router)
	http.ListenAndServe(":8080", router)
}

func configureRouter(router *mux.Router) {
	users.ConfigureRouter(router.PathPrefix("/users").Subrouter())
	home.ConfigureRouter(router.PathPrefix("/").Subrouter())
	router.NotFoundHandler = &handlers.NotFoundHandler{}
}

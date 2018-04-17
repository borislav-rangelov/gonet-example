package main

import (
	"github.com/borislav-rangelov/gonet-example/handlers/home"
	"github.com/borislav-rangelov/gonet-example/handlers/users"
	"github.com/borislav-rangelov/gonet/handlers"
	"github.com/gorilla/mux"
)

func ConfigureRouter(router *mux.Router) {
	users.ConfigureRouter(router.PathPrefix("/users").Subrouter())
	home.ConfigureRouter(router.PathPrefix("/").Subrouter())
	router.NotFoundHandler = &handlers.NotFoundHandler{}
}

package home

import "github.com/gorilla/mux"

func ConfigureRouter(router *mux.Router) {
	router.HandleFunc("/", HelloHandler).Methods("GET")
}

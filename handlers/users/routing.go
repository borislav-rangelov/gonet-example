package users

import "github.com/gorilla/mux"

func ConfigureRouter(router *mux.Router) {
	router.HandleFunc("", GetUsersPage).Methods("GET")
}

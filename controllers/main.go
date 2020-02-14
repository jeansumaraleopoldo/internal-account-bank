package controller

import (
	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/accounts", AccountGetAll).Methods("GET")
	r.HandleFunc("/accounts/{id}", AccountGetById).Methods("GET")
	r.HandleFunc("/accounts", AccountCreate).Methods("POST")

	return r
}

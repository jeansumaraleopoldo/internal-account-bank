package controller

import (
	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/accounts", AccountGetAll).Methods("GET")

	return r
}

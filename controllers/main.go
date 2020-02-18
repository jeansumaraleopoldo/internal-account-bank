package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	r := mux.NewRouter()

	r.Use(accessControlMiddleware)
	r.HandleFunc("/accounts", AccountGetAll).Methods("GET")
	r.HandleFunc("/accounts/{id}/balance", AccountBalanceGetById).Methods("GET")
	r.HandleFunc("/accounts", AccountCreate).Methods("POST")
	r.HandleFunc("/transfers", TransferGetAll).Methods("GET")
	r.HandleFunc("/transfers", TransferCreate).Methods("POST")

	return r
}

// access control and  CORS middleware
func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

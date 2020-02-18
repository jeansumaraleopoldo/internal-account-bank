package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	. "internal-account-bank/models"
	"internal-account-bank/services"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func AccountGetAll(w http.ResponseWriter, r *http.Request) {
	accounts, err := services.GetAllAccounts()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, accounts)
}

func AccountBalanceGetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	balance, err := services.GetBalanceByAccountID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, balance)
}

func AccountCreate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var account Account
	account.Created_At = time.Now()
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	account.ID = bson.NewObjectId()
	if err := services.CreateAccount(account); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, account)
}

package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	. "github.com/jeansumaraleopoldo/internal-account-bank/dao"
	. "github.com/jeansumaraleopoldo/internal-account-bank/models"
	"gopkg.in/mgo.v2/bson"
)

var dao = AccountsDAO{}

func init() {
	config := dbConfig()
	dao.Server = config[dbUrl]
	dao.Database = config[dbName]
	dao.Connect()
}

func AccountGetAll(w http.ResponseWriter, r *http.Request) {
	accounts, err := dao.GetAllAccounts()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, accounts)
}

func AccountBalanceGetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	account, err := dao.GetBalanceByAccountID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Account ID")
		return
	}
	respondWithJson(w, http.StatusOK, account.Balance)
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
	if err := dao.CreateAccount(account); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, account)
}

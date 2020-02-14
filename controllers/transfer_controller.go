package controller

import (
	"encoding/json"
	"net/http"
	"time"

	. "github.com/jeansumaraleopoldo/internal-account-bank/dao"
	. "github.com/jeansumaraleopoldo/internal-account-bank/models"
	"gopkg.in/mgo.v2/bson"
)

var daoTransfer = AccountsDAO{}

func init() {
	config := dbConfig()
	dao.Server = config[dbUrl]
	dao.Database = config[dbName]
	dao.Connect()
}

func TransferGetAll(w http.ResponseWriter, r *http.Request) {
	accounts, err := dao.GetAllAccount()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, accounts)
}

func TransferCreate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var transfer Transfer
	transfer.Created_At = time.Now()
	if err := json.NewDecoder(r.Body).Decode(&transfer); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	transfer.ID = bson.NewObjectId()
	if err := dao.CreateAccount(transfer); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, account)
}

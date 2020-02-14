package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/jeansumaraleopoldo/internal-account-bank/dao"
	. "github.com/jeansumaraleopoldo/internal-account-bank/models"
	"gopkg.in/mgo.v2/bson"
)

var dao = AccountsDAO{}

const (
	dbUrl  = "DB_URL"
	dbName = "DB_NAME"
)

func init() {
	config := dbConfig()
	dao.Server = config[dbUrl]
	dao.Database = config[dbName]
	dao.Connect()
}

func dbConfig() map[string]string {
	conf := make(map[string]string)

	conf[dbUrl] = "localhost"
	conf[dbName] = "bank"
	return conf
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

func AccountGetAll(w http.ResponseWriter, r *http.Request) {
	accounts, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, accounts)
}

func AccountGetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	product, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}
	respondWithJson(w, http.StatusOK, product)
}

func AccountCreate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	account.ID = bson.NewObjectId()
	if err := dao.Create(account); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, account)
}
